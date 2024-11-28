package models

import "gorm.io/gorm"

type AccountType struct {
	ID          string
	Name        string
	Description string
}

type Account struct {
	Model
	UserID        string
	AccountTypeID string
	AccountType   *AccountType
	CurrencyID    uint `gorm:"<-:create"`
	Currency      *Currency
	Name          string
	Priority      int
	IsDefault     bool
	IsEnable      bool
}

func (m *Account) BeforeCreate(tx *gorm.DB) error {
	m.ID = Next()
	return nil
}

type CreateAccount struct {
	UserID        string
	AccountTypeID string
	CurrencyID    uint
	Name          string
}

type AccountResponse struct {
	ID
	UserID      string
	AccountType *AccountType
	Currency    *Currency
	Name        string
	Priority    int
	IsDefault   bool
	IsEnable    bool
}

func (a *Account) ToResponse() *AccountResponse {
	return &AccountResponse{
		ID:          a.ID,
		UserID:      a.UserID,
		AccountType: a.AccountType,
		Currency:    a.Currency,
		Name:        a.Name,
		Priority:    a.Priority,
		IsDefault:   a.IsDefault,
		IsEnable:    a.IsEnable,
	}
}
