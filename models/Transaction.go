package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	ID                   string `json:"id" gorm:"primaryKey;not null;type: text"`
	AccountID            string `json:"accountID" gorm:"type:string;not null"`
	Account              Account
	Amount               float32    `json:"amount" gorm:"notnull;type: float"`
	Date                 time.Time  `json:"date" gorm:"notnull"`
	Memo                 string     `json:"memo" gorm:"type:text"`
	IsIncome             bool       `json:"isIncome" gorm:"type:bool;default: false"`
	IncomeEffectiveMonth time.Month `json:"incomeEffectiveMonth"`
	IsCleared            bool       `json:"isCleared" gorm:"notnull;default:false"`
	IsReconciled         bool       `json:"isReconciled" gorm:"notnull;default:false"`
}

func (transaction *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	transaction.ID = uuid.NewV4().String()
	return
}
