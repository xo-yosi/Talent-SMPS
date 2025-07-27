package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/xo-yosi/Talent-SMPS/internal/app/models"
	"github.com/xo-yosi/Talent-SMPS/internal/app/repository"
	"github.com/xo-yosi/Talent-SMPS/internal/utils"
)

type StudentService struct {
	srepo repository.StudentRepository
}

func NewStudentService(srepo repository.StudentRepository) *StudentService {
	return &StudentService{srepo: srepo}
}

func (s *StudentService) RegisterStudent(studentData models.StudentRegisterRequest) error {
	lastID, err := s.srepo.GetLastCoordinatorID()
	if err != nil {
		return errors.New("failed to get last Student ID")
	}

	fmt.Println("Last Student ID:", lastID)

	newID := utils.GenerateNextCoordinatorID(lastID)

	student := models.Student{
		StudentID:   newID,
		Name:        studentData.Name,
		Age:         studentData.Age,
		PhoneNumber: studentData.PhoneNumber,
		Gender:      studentData.Gender,
		CreatedAt:   time.Now(),
	}

	err = s.srepo.CreateStudent(&student)
	if err != nil {
		return errors.New("failed to create student")
	}

	return nil
}
