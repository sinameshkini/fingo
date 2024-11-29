package models

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	Model
	UserID      string
	OrderID     string
	Type        TransactionType
	Amount      Amount
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

func (m *Transaction) ToResponse(userID string) (resp *TransferResponse) {
	for _, d := range m.Documents {
		if d.Account != nil && d.Account.UserID == userID {
			return &TransferResponse{
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
	m.ID = Next()
	return nil
}

type TransferRequest struct {
	UserID          string          `json:"user_id"`
	Type            TransactionType `json:"type"`
	OrderID         string          `json:"order_id"`
	SkipLock        bool            `json:"skip_lock"`
	DebitAccountID  string          `json:"debit_account_id"`
	CreditAccountID string          `json:"credit_account_id"`
	FeeAccountID    string          `json:"fee_account_id"`
	RawAmount       Amount          `json:"raw_amount"`
	FeeAmount       Amount          `json:"fee_amount"`
	TotalAmount     Amount          `json:"total_amount"`
	Description     string          `json:"description"`
	FeeDescription  string          `json:"fee_description"`
}

type TransferResponse struct {
	CreatedAt       time.Time       `json:"created_at"`
	TransactionID   string          `json:"transaction_id"`
	OrderID         string          `json:"order_id"`
	TransactionType TransactionType `json:"transaction_type"`
	DocumentType    DocumentType    `json:"document_type"`
	Description     string          `json:"description"`
	Comment         string          `json:"comment"`
	Amount          Amount          `json:"amount"`
	Balance         Amount          `json:"balance"`
}

type ReverseRequest struct {
	TransactionID string `json:"transaction_id"`
	UserID        string `json:"user_id"`
}
