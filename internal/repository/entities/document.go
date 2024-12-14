package entities

import (
	"github.com/sinameshkini/fingo/pkg/endpoint"
	"github.com/sinameshkini/fingo/pkg/enums"
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
	Amount         models.Amount
	Balance        models.Amount
}

func (d *Document) ToResponse(userID string) (resp *endpoint.TransactionResponse) {
	if d.Account != nil && d.Account.UserID == userID {
		return &endpoint.TransactionResponse{
			CreatedAt:       d.CreatedAt,
			TransactionID:   d.TransactionID.String(),
			OrderID:         d.Transaction.OrderID,
			TransactionType: d.Transaction.Type,
			DocumentType:    d.Type,
			Description:     d.Transaction.Description,
			Comment:         d.Comment,
			Amount:          d.Amount,
			Balance:         d.Balance,
		}
	}

	return nil
}

func (d *Document) BeforeCreate(_ *gorm.DB) error {
	d.ID = models.NextSID()
	return nil
}
