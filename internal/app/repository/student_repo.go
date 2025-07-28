package repository

import (
	"github.com/xo-yosi/Talent-SMPS/internal/app/models"
)

type StudentRepository interface {
	CreateStudent(student *models.Student) error
	GetLastCoordinatorID() (int, error)
	GetStudentWithPhoneNumber(phoneNumber string) *models.Student
	GetStudentWithStudentID(studentID int) (*models.Student, error)
	UpdateSingleMeal(studentID int, meal string) error
	UpdateMealPreferences(studentID int, breakfast, lunch, dinner bool) error
	LogMealStatus(studentID int, meal string) error
	ResetAllMeals() error
}
