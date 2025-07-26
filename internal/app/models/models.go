package models

import "time"

type Student struct {
	ID        uint   `gorm:"primaryKey"`
	StudentID string `gorm:"uniqueIndex;not null"` // e.g., "1001"
	Name      string `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	Meals     []Meal 
}

type Meal struct {
    ID        uint           `gorm:"primaryKey"`
    StudentID uint           `gorm:"not null"`
    MealType  string         `gorm:"not null"`
    Date      time.Time      `gorm:"type:date;not null"`
    CreatedAt time.Time      `gorm:"not null"`
    Student   Student        `gorm:"foreignKey:StudentID"`
}