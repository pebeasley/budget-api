package services

import (
	"github.com/pebeasley/leaves-api/database"
	"github.com/pebeasley/leaves-api/models"
)

func DeleteBudget(budgetID string) error {
	err := database.DB.Delete(&models.Budget{}, budgetID)
	return err.Error
}

func GetBudgetsByUserId(userID string) ([]models.Budget, error) {
	var budgets []models.Budget

	result := database.DB.Model(
		models.Budget{}).
		Where(models.Budget{
			UserID: userID}).
		Select("id", "name").
		Find(&budgets)
	return budgets, result.Error
}
