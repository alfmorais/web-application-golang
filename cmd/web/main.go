package main

import (
	"log"
	"myapp/pkg/config"
	"myapp/pkg/handlers"
	"myapp/pkg/render"
	"net/http"
)

const portNumber string = ":8080"

func main() {
	var app config.AppConfig

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = templateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", repo.Home)
	http.HandleFunc("/about", repo.About)

	log.Printf("Starting application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
