package endpoint

import (
	"github.com/sinameshkini/fingo/pkg/enums"
	"github.com/sinameshkini/fingo/pkg/types"
	"time"
)

type TransferRequest struct {
	UserID          string                `json:"user_id"`
	Type            enums.TransactionType `json:"type"`
	OrderID         string                `json:"order_id"`
	SkipLock        bool                  `json:"skip_lock"`
	DebitAccountID  string                `json:"debit_account_id"`
	CreditAccountID string                `json:"credit_account_id"`
	FeeAccountID    string                `json:"fee_account_id"`
	RawAmount       types.Amount          `json:"raw_amount"`
	FeeAmount       types.Amount          `json:"fee_amount"`
	TotalAmount     types.Amount          `json:"total_amount"`
	Description     string                `json:"description"`
	FeeDescription  string                `json:"fee_description"`
}

type TransferResponse struct {
	CreatedAt       time.Time             `json:"created_at"`
	TransactionID   string                `json:"transaction_id"`
	OrderID         string                `json:"order_id"`
	TransactionType enums.TransactionType `json:"transaction_type"`
	DocumentType    enums.DocumentType    `json:"document_type"`
	Description     string                `json:"description"`
	Comment         string                `json:"comment"`
	Amount          types.Amount          `json:"amount"`
	Balance         types.Amount          `json:"balance"`
}

type ReverseRequest struct {
	TransactionID string `json:"transaction_id"`
	UserID        string `json:"user_id"`
}