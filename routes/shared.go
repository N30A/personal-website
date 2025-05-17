package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type pageData struct {
	Title   string
	Year    int
	Content any
}

func newPageData(title string, content any) pageData {
	return pageData{
		Title:   title,
		Year:    time.Now().Year(),
		Content: content,
	}
}

func renderTemplate(w http.ResponseWriter, templateName string, data any) {
	const directory = "templates"

	templates, err := template.ParseFiles(
		fmt.Sprintf("%s/layout.html", directory),
		fmt.Sprintf("%s/%s.html", directory, templateName),
	)
	if err != nil {
		http.Error(w, "template error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	if err := templates.Execute(w, data); err != nil {
		http.Error(w, "render error", http.StatusInternalServerError)
		log.Println(err)
	}
}
