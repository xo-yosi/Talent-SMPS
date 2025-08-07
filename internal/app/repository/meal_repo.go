package repository

import (
	"time"

	"github.com/xo-yosi/Talent-SMPS/internal/app/models"
)

type MealRepository interface {
	GetMealAnalytics(fromDate time.Time) ([]models.MealSummary, error)
	ResetAllFalseMeals() error
}
