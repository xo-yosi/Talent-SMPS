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
