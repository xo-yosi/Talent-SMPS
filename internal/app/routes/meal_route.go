package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xo-yosi/Talent-SMPS/internal/app/handler"
	"github.com/xo-yosi/Talent-SMPS/internal/app/middleware"
)

func SetupMealRoutes(r *gin.Engine, h *handler.MealHandler) {
	protected := r.Group("/api/v1")
	{
		protected.Use(middleware.Auth)
		protected.GET("/meal-analytics", h.GetMealAnalytics)
	}
}
