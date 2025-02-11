package api

import (
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

	return router
}

func (a Api) handleInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("GET /")
	w.WriteHeader(http.StatusNoContent)
}
