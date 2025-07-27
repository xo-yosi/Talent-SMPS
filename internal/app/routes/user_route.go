package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xo-yosi/Talent-SMPS/internal/app/handler"
)

func SetupUserRoutes(r *gin.Engine, h *handler.UserHandler) {
	r.POST("/api/v1/login", h.HandlerUserLogin)
}
