package main

import (
	"fmt"
	"log"
	"net/http"

	"groupie-tracker/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	// http.HandleFunc is a function that tells your Go web server: "When someone visits THIS url, run THAT function"
	http.HandleFunc("/artist",handlers.ArtistHandler)
	fmt.Println("Starting server on http://localhost:8080")
	// opena a TCP socket start a infinite loop to listen for requests for each connection read the request process it and send back a response
	// nil means use default server mux
	// static
	http.Handle("/static/",
        http.StripPrefix("/static/",
            http.FileServer(http.Dir("static"))))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
