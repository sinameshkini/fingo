package entities

import (
	"github.com/sinameshkini/fingo/pkg/enums"
	"github.com/sinameshkini/fingo/pkg/types"
	"github.com/sinameshkini/microkit/models"
	"gorm.io/gorm"
)

type Document struct {
	models.ModelSID
	TransactionID  models.SID
	Transaction    *Transaction
	AccountID      models.SID
	Account        *Account
	AccountPartyID models.SID
	Type           enums.DocumentType
	Comment        string
	Amount         types.Amount
	Balance        types.Amount
}

func (m *Document) BeforeCreate(tx *gorm.DB) error {
	m.ID = models.Next()
	return nil
}
