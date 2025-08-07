package postgres

import (
	"time"

	"github.com/xo-yosi/Talent-SMPS/internal/app/models"
	"gorm.io/gorm"
)

type MealPostgres struct {
	db *gorm.DB
}

func NewMealPostgres(db *gorm.DB) *MealPostgres {
	return &MealPostgres{db: db}
}

func (m *MealPostgres) GetMealAnalytics(fromDate time.Time) ([]models.MealSummary, error) {
	var result []models.MealSummary
	err := m.db.Raw(`
		SELECT meal_type, COUNT(*) as total
		FROM meal_logs
		WHERE created_at >= ?
		GROUP BY meal_type
	`, fromDate).Scan(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m *MealPostgres) ResetAllFalseMeals() error {
	return m.db.Session(&gorm.Session{AllowGlobalUpdate: true}).
		Model(&models.Student{}).
		Updates(map[string]interface{}{
			"breakfast": false,
			"lunch":     false,
			"dinner":    false,
		}).Error
}
