package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type Policy struct {
	Model
	EntityType string
	EntityID   string
	Settings   SettingsP `gorm:"type:jsonb;default:null"`
	Priority   int
	IsEnable   bool
}

func (m *Policy) BeforeCreate(tx *gorm.DB) error {
	m.ID = Next()
	return nil
}

type SettingsP struct {
	Limits               *LimitsP
	Codes                map[ProcessCode]CodeP
	DefaultAccountTypeID *string
}

type LimitsP struct {
	MinBalance       *Amount
	MaxBalance       *Amount
	NumberOfAccounts map[string]uint
}

type CodeP struct {
	FeeType                 *FeeType
	FeeValue                *Amount
	MinAmountPerTransaction *Amount
	MaxAmountPerTransaction *Amount
	MaxAmountPerDay         *Amount
	MaxCountPerDay          *int
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (s *SettingsP) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := SettingsP{}
	err := json.Unmarshal(bytes, &result)
	*s = result
	return err
}

// Value return json value, implement driver.Valuer interface
func (s *SettingsP) Value() (driver.Value, error) {
	//if len(s) == 0 {
	//	return nil, nil
	//}
	return json.Marshal(s)
}

type FeeType string

const (
	FeeActual     FeeType = "actual"
	FeePercentage FeeType = "percentage"
)
