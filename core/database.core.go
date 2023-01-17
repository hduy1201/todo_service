package core

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"itss.edu.vn/todo_service/models"
)

type Database struct {
	*gorm.DB
}

func NewDatabase() (*Database, error) {
	db, err := gorm.Open(sqlite.Open(os.Getenv("DB")), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.Task{}); err != nil {
		return nil, err
	}

	return &Database{DB: db}, nil
}
