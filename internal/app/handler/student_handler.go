package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xo-yosi/Talent-SMPS/internal/app/models"
	"github.com/xo-yosi/Talent-SMPS/internal/app/repository"
	"github.com/xo-yosi/Talent-SMPS/internal/app/services"
)

type StudentHandler struct {
	service *services.StudentService
	srepo   repository.StudentRepository
}

func NewStudentHandler(s *services.StudentService, r repository.StudentRepository) *StudentHandler {
	return &StudentHandler{service: s, srepo: r}
}

func (h *StudentHandler) HandlerStudentRegister(c *gin.Context) {
	var studentRegister models.StudentRegisterRequest

	if err := c.ShouldBindJSON(&studentRegister); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	student := h.srepo.GetStudentWithPhoneNumber(studentRegister.PhoneNumber)
	if student != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Student with this phone number already exists"})
		return
	}

	err := h.service.RegisterStudent(studentRegister)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student registered successfully"})
}
func (h *StudentHandler) HandlerGetStudentByID(c *gin.Context) {
	studentID := c.Param("studentID")
	if studentID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Student ID is required"})
		return
	}

	// Convert studentID to int
	var id int
	if _, err := fmt.Sscanf(studentID, "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Student ID format"})
		return
	}

	student, err := h.srepo.GetStudentWithStudentID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve student"})
		return
	}

	if student == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	if student.Breakfast && student.Lunch && student.Dinner {
		err := h.srepo.UpdateMealPreferences(student.StudentID, false, false, false)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reset meal preferences"})
			return
		}
		student.Breakfast = false
		student.Lunch = false
		student.Dinner = false
	}

	c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) HandlerStudentMeal(c *gin.Context) {
	var mealUpdate models.MealUpdateRequest
	if err := c.ShouldBindJSON(&mealUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	student, err := h.srepo.GetStudentWithStudentID(mealUpdate.StudentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve student"})
		return
	}
	if student == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	trueCount := 0
	var mealToUpdate string
	var alreadyMarked bool

	if mealUpdate.Breakfast {
		trueCount++
		mealToUpdate = "breakfast"
		alreadyMarked = student.Breakfast
	}
	if mealUpdate.Lunch {
		trueCount++
		mealToUpdate = "lunch"
		alreadyMarked = student.Lunch
	}
	if mealUpdate.Dinner {
		trueCount++
		mealToUpdate = "dinner"
		alreadyMarked = student.Dinner
	}

	if trueCount != 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only one meal can be marked at a time"})
		return
	}

	if alreadyMarked {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("You have already eaten %s", mealToUpdate)})
		return
	}

	if err := h.srepo.UpdateSingleMeal(student.StudentID, mealToUpdate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update meal status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%s marked successfully", strings.Title(mealToUpdate))})
}
