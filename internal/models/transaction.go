package models

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	Model
	OrderID     string
	Type        TransactionType
	Amount      Amount
	Description string
	CreatedAt   time.Time
	Documents   []Document
	//ReferenceID string
}

func (m *Transaction) BeforeCreate(tx *gorm.DB) error {
	m.ID = Next()
	return nil
}
