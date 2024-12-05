package entities

import (
	"github.com/sinameshkini/fingo/pkg/endpoint"
	"github.com/sinameshkini/fingo/pkg/enums"
	"github.com/sinameshkini/fingo/pkg/types"
	"github.com/sinameshkini/microkit/models"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	models.ModelSID
	UserID      string
	OrderID     string
	Type        enums.TransactionType
	Amount      types.Amount
	Description string
	CreatedAt   time.Time
	Documents   []Document
	//ReferenceID string
}

func (m *Transaction) CheckUserID(userID string) bool {
	for _, d := range m.Documents {
		if d.Account != nil && d.Account.UserID == userID {
			return true
		}
	}

	return false
}

func (m *Transaction) ToResponse(userID string) (resp *endpoint.TransferResponse) {
	for _, d := range m.Documents {
		if d.Account != nil && d.Account.UserID == userID {
			return &endpoint.TransferResponse{
				CreatedAt:       d.CreatedAt,
				TransactionID:   m.ID.String(),
				OrderID:         m.OrderID,
				TransactionType: m.Type,
				DocumentType:    d.Type,
				Description:     m.Description,
				Comment:         d.Comment,
				Amount:          d.Amount,
				Balance:         d.Balance,
			}
		}
	}

	return nil
}

func (m *Transaction) BeforeCreate(_ *gorm.DB) error {
	m.ID = models.Next()
	return nil
}
