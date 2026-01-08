package main

import (
	"fmt"
	"log"
	"net/http"

	"groupie-tracker/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/artist",handlers.ArtistHandler)
	fmt.Println("Starting server on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
