package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xo-yosi/Talent-SMPS/internal/app/models"
)

func HandlerUserLogin(c *gin.Context) {
	// Implement user login logic here
	var UserLogin models.UserLoginRequest

	if err := c.ShouldBindJSON(&UserLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}
