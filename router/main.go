package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/app", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Application application router")
	},
	)
	log.Print("Up and ready")
	http.ListenAndServe(":8080", nil)
}
