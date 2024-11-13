package db

import "fmt"

func (db DB) CreateAd(args NewAdParams) error {
	result := db.Client.Table("advertisements").Create(args)
	if result.Error != nil {
		return fmt.Errorf("failed to create new ad: %v", result.Error)
	}

	return nil
}
