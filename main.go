package main

import (
	"DungeonManagerAPI/campaigns"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
    log.Println("Starting server...")
    r := chi.NewRouter()

    r.Use(middleware.Logger)

    r.Mount("/campaigns", campaigns.GetCampaignRouter())

    // Start the server
    log.Println("Server started on port 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
