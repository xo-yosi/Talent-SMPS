package repository

import "github.com/xo-yosi/Talent-SMPS/internal/app/models"

type UserRepository interface {
	FindUserByUserName(username string) (*models.Users, error)
}
