package models

import "gorm.io/gorm"

type Document struct {
	Model
	TransactionID  ID
	Transaction    *Transaction
	AccountID      ID
	Account        *Account
	AccountPartyID ID
	Type           DocumentType
	Comment        string
	Amount         Amount
	Balance        Amount
}

func (m *Document) BeforeCreate(tx *gorm.DB) error {
	m.ID = Next()
	return nil
}
