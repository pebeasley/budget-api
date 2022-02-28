package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Subcategory struct {
	ID         string `json:"id" gorm:"primaryKey;not null;type: text"`
	CategoryId string
	Category   Category
	Name       string
}

func (subcategory *Subcategory) BeforeCreate(tx *gorm.DB) (err error) {
	subcategory.ID = uuid.NewV4().String()
	return
}
