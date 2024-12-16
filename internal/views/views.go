package views

import (
	"html/template"
	"log"
	"net/http"
)

func IndexView(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "index.html", nil)
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmplPath := "templates/" + tmpl
	t, err := template.ParseFiles("templates/base.html", tmplPath)
	log.Printf("Rendering template: %s", tmplPath)

	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, "Template exectuion error:"+err.Error(), http.StatusInternalServerError)
		log.Printf("Error executing templates %v", err)
		return
	}
}
