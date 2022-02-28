package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Budget struct {
	ID           string `json:"id" gorm:"primaryKey;not null;type: text"`
	UserID       string
	User         User
	Name         string
	BudgetConfig *BudgetConfig
	Accouts      []Account
	Categories   []Category
	BudgetMonths []BudgetMonth
}

func (budget *Budget) BeforeCreate(tx *gorm.DB) (err error) {
	budget.ID = uuid.NewV4().String()
	return
}
