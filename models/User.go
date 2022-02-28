package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string    `json:"id" gorm:"primaryKey;not null;type: text"`
	CreatedAt time.Time `json:"created_at" gorm:"type:time;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:time;not null"`
	FirstName string    `json:"firstName" gorm:"type:text;not null"`
	LastName  string    `json:"lastName" gorm:"type:text;not null"`
	Email     string    `json:"email" gorm:"uniqueIndex;unique;type:text;not null"`
	Password  string    `json:"password" gorm:"index;type:text;not null"`
	IsActive  bool      `json:"isActive" gorm:"type:bool;not null;default:true"`
	Budgets   []Budget
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.NewV4().String()
	return
}
