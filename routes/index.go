package routes

import "net/http"

func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", newPageData("Hem", "index", nil))
}
