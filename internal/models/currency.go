package models

import (
	"encoding/json"
)

type Currency struct {
	ID        uint `gorm:"primarykey"`
	Symbol    string
	Icon      string
	IsDefault bool
	IsEnable  bool
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
