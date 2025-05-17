package handlers

import "net/http"

type indexData struct {
	Message string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	indexData := indexData{
		Message: "Hello World",
	}
	renderTemplate(w, "index", newPageData("Hem", indexData))
}
