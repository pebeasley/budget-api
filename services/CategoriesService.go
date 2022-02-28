package services

import (
	"github.com/pebeasley/leaves-api/database"
	"github.com/pebeasley/leaves-api/models"
)

func GetCategoriesByBudgetID(budgetID string) ([]models.Category, error) {
	var categories []models.Category

	result := database.DB.Model(models.Category{}).
		Where("budget_id = ?", budgetID).
		Find(&categories)

	return categories, result.Error
}
