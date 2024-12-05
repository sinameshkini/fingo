package endpoint

import "github.com/sinameshkini/fingo/pkg/types"

type AccountType struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
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
	Balance     types.Amount `json:"balance"`
}

type Currency struct {
	ID        uint   `json:"id"`
	Symbol    string `json:"symbol"`
	Icon      string `json:"icon"`
	IsDefault bool   `json:"is_default"`
	IsEnable  bool   `json:"is_enable"`
}
