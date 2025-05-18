package routes

import "net/http"

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about", newPageData("Om mig", "about", nil))
}
