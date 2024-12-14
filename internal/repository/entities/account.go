package entities

import (
	"github.com/sinameshkini/fingo/pkg/endpoint"
	"github.com/sinameshkini/fingo/pkg/enums"
	"github.com/sinameshkini/microkit/models"
	"gorm.io/gorm"
)

type AccountType struct {
	ID          string `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (m *AccountType) ToResponse() *endpoint.AccountType {
	return &endpoint.AccountType{
		ID:          m.ID,
		Name:        m.Name,
		Description: m.Description,
	}
}

type Account struct {
	models.ModelSID
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
		return nil, enums.ErrNotFound
	}

	return m.settings, nil
}

func (m *Account) BeforeCreate(_ *gorm.DB) error {
	m.ID = models.NextSID()
	return nil
}

func (a *Account) ToResponse(balance models.Amount) *endpoint.AccountResponse {
	return &endpoint.AccountResponse{
		ID:          a.ID.String(),
		UserID:      a.UserID,
		AccountType: a.AccountType.ToResponse(),
		Currency:    a.Currency.ToResponse(),
		Name:        a.Name,
		Priority:    a.Priority,
		IsDefault:   a.IsDefault,
		IsEnable:    a.IsEnable,
		Balance:     balance,
	}
}
