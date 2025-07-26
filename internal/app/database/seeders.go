package database

import (
	"time"

	"github.com/xo-yosi/Talent-SMPS/internal/app/models"
	"gorm.io/gorm"
)

func Seeders(db *gorm.DB) error {
	users := []models.Users{
		{
			Username:  "admin1",
			Password:  "admin123",
			CreatedAt: time.Now(),
		},
		{
			Username:  "admin2",
			Password:  "admin123",
			CreatedAt: time.Now(),
		},
		{
			Username:  "admin3",
			Password:  "admin123",
			CreatedAt: time.Now(),
		},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			return err
		}
	}

	return nil
}
