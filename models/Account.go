package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Account struct {
	ID           string `json:"id" gorm:"primaryKey;not null;type: text"`
	BudgetId     string
	Budget       Budget
	Transactions []Transaction
	Name         string
	Type         string
}

func (account *Account) BeforeCreate(tx *gorm.DB) (err error) {
	account.ID = uuid.NewV4().String()
	return
}
