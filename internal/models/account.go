package models

import "gorm.io/gorm"

type AccountType struct {
	ID          string `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
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
	settings      *Settings `gorm:"-"`
}

func (m *Account) SetSettings(settings *Settings) {
	m.settings = settings
}

func (m *Account) GetSettings() (*Settings, error) {
	if m.settings == nil {
		return nil, ErrNotFound
	}

	return m.settings, nil
}

func (m *Account) BeforeCreate(_ *gorm.DB) error {
	m.ID = Next()
	return nil
}

type CreateAccount struct {
	UserID        string `json:"user_id"`
	AccountTypeID string `json:"account_type_id"`
	CurrencyID    uint   `json:"currency_id"`
	Name          string `json:"name"`
}

type AccountResponse struct {
	ID          string       `json:"id"`
	UserID      string       `json:"user_id"`
	AccountType *AccountType `json:"account_type"`
	Currency    *Currency    `json:"currency"`
	Name        string       `json:"name"`
	Priority    int          `json:"priority"`
	IsDefault   bool         `json:"is_default"`
	IsEnable    bool         `json:"is_enable"`
	Balance     Amount       `json:"balance"`
}

func (a *Account) ToResponse(balance Amount) *AccountResponse {
	return &AccountResponse{
		ID:          a.ID.String(),
		UserID:      a.UserID,
		AccountType: a.AccountType,
		Currency:    a.Currency,
		Name:        a.Name,
		Priority:    a.Priority,
		IsDefault:   a.IsDefault,
		IsEnable:    a.IsEnable,
		Balance:     balance,
	}
}
