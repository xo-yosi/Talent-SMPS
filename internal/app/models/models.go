package models

import (
	"time"

	"github.com/google/uuid"
)

type Student struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	StudentID   int       `json:"student_id" gorm:"uniqueIndex;not null"`
	Name        string    `json:"name" gorm:"not null"`
	Age         int       `json:"age" gorm:"not null"`
	PhoneNumber string    `json:"phone_number" gorm:"not null"`
	Gender      string    `json:"gender" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null"`
	Breakfast   bool      `json:"breakfast" gorm:"not null"`
	Lunch       bool      `json:"lunch" gorm:"not null"`
	Dinner      bool      `json:"dinner" gorm:"not null"`
}

// type MealLog struct {
// 	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
// 	StudentID uuid.UUID `json:"student_id" gorm:"not null"`
// 	MealType  string    `json:"meal_type" gorm:"not null"`
// 	CreatedAt time.Time `json:"created_at" gorm:"not null"`
// 	Student   Student   `gorm:"foreignKey:StudentID;references:ID"`
// }
