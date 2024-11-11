package db

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	Client *gorm.DB
}

func NewClient() (*DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return &DB{}, fmt.Errorf("failed to create database client: %v", err)
	}

	return &DB{db}, nil
}
