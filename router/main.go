package main

import (
	"log"
	"net/http"
)

func main() {
	log.Print("Starting router...")
	router := Router()
	log.Print("Starting web server")
	log.Fatal(http.ListenAndServe(":8080", router))
}
