package main

import (
	"log"
	"net/http"

	"./handler"
)

func main() {
	log.Print("Starting router...")
	router := handler.Router()
	log.Print("Starting web server")
	log.Fatal(http.ListenAndServe(":8080", router))
}
