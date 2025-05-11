package main

import (
	"censor/pkg/api"
	"censor/pkg/censor"
	"log"
	"net/http"
)

func main() {
	// Запускаем API
	api := api.New(censor.New())
	log.Print("Server is starting...")
	http.ListenAndServe(":8080", api.Router())
	log.Print("Server has been stopped.")
}
