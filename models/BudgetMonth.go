package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type BudgetMonth struct {
	ID       string `json:"id" gorm:"primaryKey;not null;type: text"`
	Year     int16
	Month    int16
	BudgetID string
	Budget   Budget
}

func (budgetMonth *BudgetMonth) BeforeCreate(tx *gorm.DB) (err error) {
	budgetMonth.ID = uuid.NewV4().String()
	return
}
