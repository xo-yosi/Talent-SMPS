package repository

import "github.com/xo-yosi/Talent-SMPS/internal/app/models"

type StudentRepository interface {
	CreateStudent(student *models.Student) error
	GetLastCoordinatorID() (int, error)
}
