package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	Client *gorm.DB
}

func NewClient(dsn string) (*DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return &DB{}, fmt.Errorf("failed to create database client: %v", err)
	}

	return &DB{db}, nil
}
