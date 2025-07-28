package postgres

import (
	"errors"

	"github.com/xo-yosi/Talent-SMPS/internal/app/models"
	"gorm.io/gorm"
)

type StudentPostgres struct {
	db *gorm.DB
}

func NewStudentPostgres(db *gorm.DB) *StudentPostgres {
	return &StudentPostgres{db: db}
}

func (s *StudentPostgres) CreateStudent(student *models.Student) error {
	if err := s.db.Create(&student).Error; err != nil {
		return err
	}
	return nil
}

func (s *StudentPostgres) GetLastCoordinatorID() (int, error) {
	var student models.Student

	err := s.db.Order("created_at desc").First(&student).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return student.StudentID, nil
}

func (s *StudentPostgres) GetStudentWithPhoneNumber(phoneNumber string) *models.Student {
	var student models.Student
	err := s.db.Where("phone_number = ?", phoneNumber).First(&student).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	if err != nil {
		return nil
	}
	return &student
}

func (s *StudentPostgres) GetStudentWithStudentID(studentID int) (*models.Student, error) {
	var student models.Student
	err := s.db.Where("student_id = ?", studentID).First(&student).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *StudentPostgres) UpdateSingleMeal(studentID int, meal string) error {
	updates := map[string]interface{}{meal: true}
	return r.db.Model(&models.Student{}).Where("student_id = ?", studentID).Updates(updates).Error
}

func (r *StudentPostgres) UpdateMealPreferences(studentID int, breakfast, lunch, dinner bool) error {
	return r.db.Model(&models.Student{}).Where("student_id = ?", studentID).
		Updates(map[string]interface{}{
			"breakfast": breakfast,
			"lunch":     lunch,
			"dinner":    dinner,
		}).Error
}