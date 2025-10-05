package main

import (
	"fmt"
	"log"
	"net/http"
)

const portNumber string = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("This is the home page"))
	if err != nil {
		log.Println("Error writing response:", err)
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := AddValues(3, 5)
	_, _ = fmt.Fprintf(w, "This is the about page. The sum of 3 and 5 is %d", sum)
}

func AddValues(x, y int) int {
	var sum int = x + y
	return sum
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Printf("Starting application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
