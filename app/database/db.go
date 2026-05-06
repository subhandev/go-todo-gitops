package database

import (
	"fmt"
	"os"

	"github.com/subhandev/go-todo-gitops/app/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	path := os.Getenv("SQLITE_PATH")
	if path == "" {
		path = "todos.db"
	}

	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("open sqlite db: %w", err)
	}

	if err := db.AutoMigrate(&models.Todo{}); err != nil {
		return nil, fmt.Errorf("auto-migrate: %w", err)
	}

	return db, nil
}

