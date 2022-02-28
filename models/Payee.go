package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Payee struct {
	ID   string `json:"id" gorm:"primaryKey;not null;type: text"`
	Name string
}

func (payee *Payee) BeforeCreate(tx *gorm.DB) (err error) {
	payee.ID = uuid.NewV4().String()
	return
}
