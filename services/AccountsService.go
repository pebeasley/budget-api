package services

import (
	"github.com/pebeasley/leaves-api/database"
	"github.com/pebeasley/leaves-api/models"
)

func GetAccountsByBudgetID(budgetID string) ([]models.Account, error) {
	var foundAccounts []models.Account
	result := database.DB.
		Model(models.Account{}).
		Where("user_id = ?", budgetID).
		Find(&foundAccounts)
	return foundAccounts, result.Error
}
