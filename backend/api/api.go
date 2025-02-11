package api

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

type Api struct {
	logger *log.Logger
	port   string
}

func New() Api {
	port := "8080"

	if p, hasPort := os.LookupEnv("PORT"); hasPort {
		port = p
	}

	return Api{
		logger: log.New(os.Stderr, "[api]", log.LstdFlags),
		port:   ":" + port,
	}
}

func (a Api) Run() {
	log.Fatal(http.ListenAndServe(a.port, a.Router()))
}

func (a Api) Router() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", a.handleInfo)
	router.GET("/:stories", a.handleStories)

	return router
}

func (a Api) handleInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("GET /")
	w.WriteHeader(http.StatusNoContent)
}

func (a Api) handleStories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println("GET /" + ps.ByName("stories"))
	stories := []string{"topstories", "newstories", "beststories"}

	if !contains(stories, ps.ByName("stories")) {
		http.Error(w, "Invalid stories type", http.StatusBadRequest)
		return
	}

	url := "https://hacker-news.firebaseio.com/v0/" + ps.ByName("stories") + ".json"

	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		http.Error(w, "Error fetching data: "+err.Error(), http.StatusInternalServerError)
		a.logger.Println(err)
		return
	}

	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading response: "+err.Error(), http.StatusInternalServerError)
		a.logger.Println(err)
		return
	}

	w.Write(body)
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
