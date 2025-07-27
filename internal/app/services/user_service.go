package services

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/xo-yosi/Talent-SMPS/internal/app/models"
	"github.com/xo-yosi/Talent-SMPS/internal/app/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	urepo repository.UserRepository
}

func NewUserService(urepo repository.UserRepository) *UserService {
	return &UserService{urepo: urepo}
}

func (s *UserService) UserLogin(input models.UserLoginRequest) (accessToken string, refreshToken string, err error) {
	user, err := s.urepo.FindUserByUserName(input.Username)
	if err != nil {
		return "", "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return "", "", errors.New("invalid password")
	}

	accessClaims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    "admin",
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	accessTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err = accessTokenObj.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", "", errors.New("failed to generate access token")
	}

	refreshClaims := jwt.MapClaims{
		"user_id": user.ID,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
		"type":    "refresh",
	}
	refreshTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err = refreshTokenObj.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", "", errors.New("failed to generate refresh token")
	}

	return accessToken, refreshToken, nil
}
