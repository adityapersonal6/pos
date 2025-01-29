package main

import (
	"log"

	"github.com/adityapersonal6/pos/internal/config"
	"github.com/adityapersonal6/pos/internal/database"
	"github.com/adityapersonal6/pos/internal/routes"
)

const filePath = "app-conf.json"

func main() {
	cfg, err := config.LoadConfig(filePath)
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	defer db.Close()

	router := routes.SetupRoutes(db)

	log.Printf("Starting server on %s", cfg.ServerURL)
	router.Run(cfg.ServerURL)

}
