package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type BudgetConfig struct {
	ID       string `json:"id" gorm:"primaryKey;not null;type: text"`
	BudgetID string
	Budget   Budget
}

func (budgetConfig *BudgetConfig) BeforeCreate(tx *gorm.DB) (err error) {
	budgetConfig.ID = uuid.NewV4().String()
	return
}
