package main

import (
	"log"
	"net"
	"net/http"

	"github.com/N30A/personal-website/db"
	"github.com/N30A/personal-website/routes"
)

const (
	Host = "0.0.0.0"
	Port = "8080"
)

func main() {
	if err := db.Connect(); err != nil {
		log.Fatalln("database connection failed:", err)
	}
	defer db.DB.Close()

	mux := http.NewServeMux()
	routes.MapStaticFiles(mux)
	routes.MapRoutes(mux)

	address := net.JoinHostPort(Host, Port)
	log.Printf("listening on %s\n", address)
	log.Fatalln(http.ListenAndServe(address, mux))
}
