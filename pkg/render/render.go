package render

import (
	"bytes"
	"html/template"
	"log"
	"myapp/pkg/config"
	"myapp/pkg/models"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}
var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmplDir string, td *models.TemplateData) {
	var templateToRender map[string]*template.Template
	var name string = filepath.Base(tmplDir)

	if app.UseCache {
		templateToRender = app.TemplateCache
	} else {
		templateToRender, _ = CreateTemplateCache()
	}

	tc, ok := templateToRender[name]
	if !ok {
		handleInternalServerError("template not found in cache: "+tmplDir, w, nil)
		return
	}

	td = AddDefaultData(td)
	buf := new(bytes.Buffer)
	err := tc.Execute(buf, td)
	if err != nil {
		handleInternalServerError("error executing template: "+tmplDir, w, err)
		return
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		handleInternalServerError("error writing template to browser: "+tmplDir, w, err)
		return
	}
}

func handleInternalServerError(message string, w http.ResponseWriter, err error) {
	log.Println(message, err)
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.go.tmpl")
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return cache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return cache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return cache, err
			}
		}
		cache[name] = ts
	}
	return cache, nil
}
