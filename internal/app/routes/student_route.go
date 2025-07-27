package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xo-yosi/Talent-SMPS/internal/app/handler"
	"github.com/xo-yosi/Talent-SMPS/internal/app/middleware"
)

func SetupStudentRoutes(r *gin.Engine, h *handler.StudentHandler) {
	protected := r.Group("/api/v1")
	{
		protected.Use(middleware.Auth)
		protected.POST("/student-register", h.HandlerStudentRegister)
		protected.GET("/student/:studentID", h.HandlerGetStudentByID)
		protected.POST("/student/update-meal", h.HandlerStudentMeal)
	}
}
