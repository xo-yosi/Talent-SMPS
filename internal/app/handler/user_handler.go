package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xo-yosi/Talent-SMPS/internal/app/models"
	"github.com/xo-yosi/Talent-SMPS/internal/app/repository"
	"github.com/xo-yosi/Talent-SMPS/internal/app/services"
)

type UserHandler struct {
	service *services.UserService
	urepo   repository.UserRepository
}

func NewUserHandler(s *services.UserService, r repository.UserRepository) *UserHandler {
	return &UserHandler{service: s, urepo: r}
}

func (h *UserHandler) HandlerUserLogin(c *gin.Context) {
	var UserLogin models.UserLoginRequest

	if err := c.ShouldBindJSON(&UserLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.urepo.FindUserByUserName(UserLogin.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	token, refreshToken, err := h.service.UserLogin(UserLogin)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	response := gin.H{
		"user_id":       user.ID,
		"username":      user.Username,
		"access_token":  token,
		"refresh_token": refreshToken,
		"message":       "Login successful",
	}

	c.JSON(http.StatusOK, response)
}
