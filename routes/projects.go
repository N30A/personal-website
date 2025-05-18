package routes

import "net/http"

type projectsData struct {
	Projects []project
}

type project struct {
	Title      string
	Body       string
	ImageUrl   string
	ProjectUrl string
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	projectsData := projectsData{
		Projects: []project{
			{
				Title:      "Webbapp för tidrapportering",
				Body:       "Ett verktyg för att logga arbetstid, byggt i Go med PostgreSQL och HTMX.",
				ImageUrl:   "https://placehold.co/300x250",
				ProjectUrl: "https://github.com/n30a",
			},
			{
				Title:      "Portfolio-webbplats",
				Body:       "En responsiv webbplats för att visa upp mina projekt. Använder Go templates och ren CSS.",
				ImageUrl:   "https://placehold.co/300x250",
				ProjectUrl: "https://github.com/n30a",
			},
			{
				Title:      "Chattapplikation i Go",
				Body:       "En realtidschatt byggd med WebSockets och Go – frontend i HTMX.",
				ImageUrl:   "https://placehold.co/300x250",
				ProjectUrl: "https://github.com/n30a",
			},
			{
				Title:      "CLI-verktyg för filhantering",
				Body:       "Ett kommandoradsverktyg i Go för att organisera filer och mappar automatiskt.",
				ImageUrl:   "https://placehold.co/300x250",
				ProjectUrl: "https://github.com/n30a",
			},
			{
				Title:      "Bloggmotor i Go",
				Body:       "En enkel bloggmotor med markdown-stöd, SQLite och inbyggd adminpanel.",
				ImageUrl:   "https://placehold.co/300x250",
				ProjectUrl: "https://github.com/n30a",
			},
			{
				Title:      "API för väderdata",
				Body:       "REST-API i Go som hämtar och cachar väderdata från externa tjänster.",
				ImageUrl:   "https://placehold.co/300x250",
				ProjectUrl: "https://github.com/n30a",
			},
		},
	}

	renderTemplate(w, "projects", newPageData("Projekt", "projects", projectsData))
}
