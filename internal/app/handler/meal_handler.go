package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xo-yosi/Talent-SMPS/internal/app/repository"
)

type MealHandler struct {
	mrepo repository.MealRepository
	srepo repository.StudentRepository
}

func NewMealHandler(m repository.MealRepository, s repository.StudentRepository) *MealHandler {
	return &MealHandler{mrepo: m, srepo: s}
}

func (h *MealHandler) GetMealAnalytics(c *gin.Context) {
	rangeParam := c.DefaultQuery("range", "today")
	var fromDate time.Time

	now := time.Now()

	switch rangeParam {
	case "today":
		fromDate = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	case "7d":
		fromDate = now.AddDate(0, 0, -7)
	case "14d":
		fromDate = now.AddDate(0, 0, -14)
	case "1m":
		fromDate = now.AddDate(0, -1, 0)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid range"})
		return
	}

	result, err := h.mrepo.GetMealAnalytics(fromDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch meal analytics"})
		return
	}

	totalStudents, err := h.srepo.GetTotalStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch total students"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_students": totalStudents,
		"result":          result,
	})
}
