package render

import (
	"html/template"
	"log"
	"net/http"
)

const baseLayoutTemplate string = "./templates/base.layout.tmpl"

func RenderTemplate(w http.ResponseWriter, tmplDir string) {
	parsedTemplate, _ := template.ParseFiles(tmplDir, baseLayoutTemplate)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("error parsing template:", err)
		http.Error(w, "Internal Server Error", 500)
	}
}
