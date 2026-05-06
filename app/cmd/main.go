package main

import (
	"log"
	"os"

	"github.com/subhandev/go-todo-gitops/app/database"
	"github.com/subhandev/go-todo-gitops/app/routes"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("db: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := routes.New(db)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("server: %v", err)
	}
}

