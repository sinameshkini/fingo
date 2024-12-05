package entities

import (
	"encoding/json"
	"github.com/sinameshkini/fingo/pkg/endpoint"
)

type Currency struct {
	ID        uint `gorm:"primarykey"`
	Symbol    string
	Icon      string
	IsDefault bool
	IsEnable  bool
}

func (m *Currency) ToResponse() *endpoint.Currency {
	return &endpoint.Currency{
		ID:        m.ID,
		Symbol:    m.Symbol,
		Icon:      m.Icon,
		IsDefault: m.IsDefault,
		IsEnable:  m.IsEnable,
	}
}

type CurrencyObject struct {
	Currencies []*Currency
}

func (m *CurrencyObject) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

func (m *CurrencyObject) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}
