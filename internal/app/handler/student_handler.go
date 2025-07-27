package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xo-yosi/Talent-SMPS/internal/app/repository"
	"github.com/xo-yosi/Talent-SMPS/internal/app/services"
	"github.com/xo-yosi/Talent-SMPS/internal/app/models"
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

	student, err := h.srepo.GetStudentWithStudentID(studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve student"})
		return
	}

	if student == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	c.JSON(http.StatusOK, student)
}
