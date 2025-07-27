package repository

import (
	"github.com/xo-yosi/Talent-SMPS/internal/app/models"
)

type StudentRepository interface {
	CreateStudent(student *models.Student) error
	GetLastCoordinatorID() (int, error)
	GetStudentWithPhoneNumber(phoneNumber string) *models.Student
	GetStudentWithStudentID(studentID int) (*models.Student, error)
	MarkMeal(studentID int, meal string) error
	UpdateSingleMeal(studentID int, meal string) error
	ResetDailyMeals(studentID int) error
}
