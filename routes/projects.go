package routes

import (
	"log"
	"net/http"

	"github.com/N30A/personal-website/db"
)

type projectsData struct {
	Projects []project
}

type technology struct {
	name string `db:"name"`
}

type project struct {
	Title        string `db:"title"`
	Body         string `db:"body"`
	ImageName    string `db:"image_name"`
	ProjectUrl   string `db:"project_url"`
	Technologies []technology
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	query := `
	select title, body, image_name, project_url from projects;
	`

	projects := []project{}
	if err := db.DB.Select(&projects, query); err != nil {
		log.Println(err)
	}

	renderTemplate(w, "projects", newPageData("Projekt", "projects", projectsData{projects}))
}
