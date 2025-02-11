package main

import (
	"log"
	"rn-reader-backend/api"
)

func main() {
	runApi()
}

func runApi() {
	log.Println("Run API")
	api := api.New()
	api.Run()
}
