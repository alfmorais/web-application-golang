package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

const portNumber string = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Request GET /")
	_, err := w.Write([]byte("This is the home page"))
	if err != nil {
		log.Println("Error writing response:", err)
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	log.Println("Request GET /about")
	sum := AddValues(3, 5)
	_, _ = fmt.Fprintf(w, "This is the about page. The sum of 3 and 5 is %d", sum)
}

func Divide(w http.ResponseWriter, r *http.Request) {
	log.Println("Request GET /divide")
	result, err := divide(10.0, 20.0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, _ = fmt.Fprintf(w, "The result of division is %.2f", result)
}

func AddValues(x, y int) int {
	var sum int = x + y
	return sum
}

func divide(x, y float32) (float32, error) {
	if y == 0 {
		return 0, errors.New("division by zero")
	}
	return x / y, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	log.Printf("Starting application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
