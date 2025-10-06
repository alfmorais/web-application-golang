package handlers

import (
	"log"
	"myapp/pkg/render"
	"net/http"
)

const templateDir string = "./templates/"

func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Request GET /")
	fullTemplatePath := templateDir + "home.page.go.tmpl"
	render.RenderTemplate(w, fullTemplatePath)
}

func About(w http.ResponseWriter, r *http.Request) {
	log.Println("Request GET /about")
	fullTemplatePath := templateDir + "about.page.go.tmpl"
	render.RenderTemplate(w, fullTemplatePath)
}
