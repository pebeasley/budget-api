package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID       string `json:"id" gorm:"primaryKey;not null;type: text"`
	BudgetID string `json:"budget_id" gorm:"type: text"`
	Budget   Budget
	Name     string `json:"name" gorm:"type: text"`
	IsHidden bool   `json:"is_hidden" gorm:"type: bool"`
}

func (category *Category) BeforeCreate(tx *gorm.DB) (err error) {
	category.ID = uuid.NewV4().String()
	return
}
