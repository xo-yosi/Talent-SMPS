package models

import (
	"time"

	"github.com/google/uuid"
)

type Student struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	StudentID   int       `json:"student_id" gorm:"uniqueIndex;not null"`
	// ProfilePic  string    `json:"profile_pic" gorm:"not null"`
	Name        string    `json:"name" gorm:"not null"`
	Age         int       `json:"age" gorm:"not null"`
	PhoneNumber string    `json:"phone_number" gorm:"not null"`
	Gender      string    `json:"gender" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null"`
	Breakfast   bool      `json:"breakfast" gorm:"not null" default:"false"`
	Lunch       bool      `json:"lunch" gorm:"not null" default:"false"`
	Dinner      bool      `json:"dinner" gorm:"not null" default:"false"`
}

type Users struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Username  string    `json:"username" gorm:"uniqueIndex;not null"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
}

type StudentRegisterRequest struct {
	Name        string `json:"name" binding:"required"`
	Age         int    `json:"age" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Gender      string `json:"gender" binding:"required"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type MealUpdateRequest struct {
	StudentID int  `json:"student_id" binding:"required"`
	Breakfast bool `json:"breakfast"`
	Lunch     bool `json:"lunch"`
	Dinner    bool `json:"dinner"`
}

type MealLog struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	StudentID int       `json:"student_id" gorm:"not null;index foreignKey:StudentID;references:StudentID"`
	MealType  string    `json:"meal_type" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
}
