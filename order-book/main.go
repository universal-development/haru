package main

import (
	"log"
	"net/http"

	"./handler"
)

func main() {
	log.Print("Starting order book application...")
	router := handler.Router()
	log.Fatal(http.ListenAndServe(":8080", router))
}
