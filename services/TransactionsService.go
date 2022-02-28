package services

import (
	"github.com/pebeasley/leaves-api/database"
	"github.com/pebeasley/leaves-api/models"
)

func GetTransaction(id string) models.Transaction {
	var FoundTransaction models.Transaction
	response := database.DB.
		Where(&models.Transaction{ID: id}).
		Find(&FoundTransaction)

	if response.Error != nil {
		panic("An error occurred when retrieving the transaction: " + response.Error.Error())
	}
	return FoundTransaction
}

func GetTransactionsByAccount(accountID string) ([]models.Transaction, error) {
	var Transactions []models.Transaction
	response := database.DB.
		Model(models.Transaction{}).
		Where("account_id = ?", accountID).
		Find(&Transactions)

	return Transactions, response.Error
}
