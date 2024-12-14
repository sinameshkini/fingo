package entities

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sinameshkini/fingo/pkg/enums"
	"github.com/sinameshkini/microkit/models"
	"gorm.io/gorm"
)

type Policy struct {
	models.ModelSID
	EntityType string
	EntityID   string
	Settings   SettingsP `gorm:"type:jsonb;default:null"`
	Priority   int
	IsEnable   bool
}

func (m *Policy) BeforeCreate(tx *gorm.DB) error {
	m.ID = models.NextSID()
	return nil
}

type SettingsP struct {
	//Limits               *LimitsP
	Limits               map[string]LimitsP
	Codes                map[enums.ProcessCode]CodeP
	DefaultAccountTypeID *string
}

type LimitsP struct {
	MinBalance       *models.Amount
	MaxBalance       *models.Amount
	NumberOfAccounts *int
}

type CodeP struct {
	FeeType                 *FeeType
	FeeValue                *models.Amount
	MinAmountPerTransaction *models.Amount
	MaxAmountPerTransaction *models.Amount
	MaxAmountPerDay         *models.Amount
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

type Settings struct {
	//Limits               Limits
	Limits               map[string]Limits
	Codes                map[enums.ProcessCode]Code
	DefaultAccountTypeID string
}

type Limits struct {
	MinBalance       models.Amount
	MaxBalance       models.Amount
	NumberOfAccounts int
}

type Code struct {
	FeeType                 FeeType
	FeeValue                models.Amount
	MinAmountPerTransaction models.Amount
	MaxAmountPerTransaction models.Amount
	MaxAmountPerDay         models.Amount
	MaxCountPerDay          int
}

func (c *Code) CalculateFeeAmount(raw models.Amount) (fee models.Amount) {
	switch c.FeeType {
	case FeeActual:
		fee = c.FeeValue
	case FeePercentage:
		fee = raw * c.FeeValue / 100
	default:
		fee = 0
	}

	return
}
