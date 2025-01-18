package db

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	Client *gorm.DB
	Cache  *cache.Cache
}

func NewClient(dsn string) (*DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return &DB{}, fmt.Errorf("failed to create database client: %v", err)
	}

	err = db.AutoMigrate(
		&Moderation{},
	)
	if err != nil {
		return &DB{}, fmt.Errorf("failed to execute migrate action: %v", err)
	}

	cache := cache.New(24*time.Hour, 24*time.Hour)

	return &DB{db, cache}, nil
}
