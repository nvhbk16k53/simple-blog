package blog

import (
	"html/template"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/nvhbk16k53/simple-blog/db"
)

var templatePath = "template/"

var templates = template.Must(template.ParseGlob(templatePath + "*.tmpl"))

/* Render index page template */
func RenderIndex(w http.ResponseWriter) {
	err := templates.ExecuteTemplate(w, "index.tmpl", nil)
	if err != nil {
		log.Errorf("Cannot execute template index.tmpl: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/* Render article page template */
func RenderArticle(w http.ResponseWriter, tmpl string, article *db.Article) {
	err := templates.ExecuteTemplate(w, tmpl + ".tmpl", article)
	if err != nil {
		log.Errorf("Cannot execute template %v.tmpl: %v", tmpl, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
