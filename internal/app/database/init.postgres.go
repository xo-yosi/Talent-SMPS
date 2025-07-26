package database

import (
	"github.com/xo-yosi/Talent-SMPS/internal/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(databaseURL string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Student{}, 
		// &models.MealLog{},
	)
	if err != nil {
		return err
	}
	return nil
}
