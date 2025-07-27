package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xo-yosi/Talent-SMPS/internal/app/handler"
)

func SetupStudentRoutes(r *gin.Engine, h *handler.StudentHandler) {
	r.POST("/api/v1/register-student", h.HandlerStudentRegister)
}
