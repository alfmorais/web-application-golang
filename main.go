package main

import (
	"log"
	"net/http"
)

const portNumber string = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Request GET /")
}

func About(w http.ResponseWriter, r *http.Request) {
	log.Println("Request GET /about")
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Printf("Starting application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
