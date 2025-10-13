package handlers

import (
	"log"
	"myapp/pkg/config"
	"myapp/pkg/models"
	"myapp/pkg/render"
	"net/http"
)

const templateDir string = "./templates/"

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Request GET /")
	fullTemplatePath := templateDir + "home.page.go.tmpl"
	render.RenderTemplate(w, fullTemplatePath, &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	log.Println("Request GET /about")
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."
	fullTemplatePath := templateDir + "about.page.go.tmpl"
	render.RenderTemplate(w, fullTemplatePath, &models.TemplateData{
		StringMap: stringMap,
	})
}
