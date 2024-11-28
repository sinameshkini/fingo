package models

import "gorm.io/gorm"

type Document struct {
	Model
	TransactionID  string
	Transaction    *Transaction
	AccountID      string
	Account        *Account
	AccountPartyID string
	Type           DocumentType
	Comment        string
	Amount         Amount
	Balance        Amount
}

func (m *Document) BeforeCreate(tx *gorm.DB) error {
	m.ID = Next()
	return nil
}
