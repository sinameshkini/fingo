package endpoint

import (
	"github.com/sinameshkini/fingo/pkg/enums"
	"github.com/sinameshkini/microkit/models"
	"time"
)

type TransactionRequest struct {
	UserID      string                `json:"user_id" validate:"required"`
	OrderID     string                `json:"order_id" validate:"required"`
	Type        enums.TransactionType `json:"type" validate:"required"`
	TotalAmount models.Amount         `json:"total_amount" validate:"gt=0"`
	Description string                `json:"description" validate:"required"`
	Transfers   []TransferRequest     `json:"transfers" validate:"required"`
}

type TransferRequest struct {
	SkipLock        bool          `json:"skip_lock"`
	DebitAccountID  string        `json:"debit_account_id" validate:"required"`
	CreditAccountID string        `json:"credit_account_id" validate:"required"`
	Amount          models.Amount `json:"amount" validate:"gt=0"`
	Comment         string        `json:"comment" validate:"required"`
}

func (r *TransactionRequest) ValidateAmount() error {
	var sumAmounts models.Amount
	for _, transfer := range r.Transfers {
		sumAmounts += transfer.Amount
	}

	if sumAmounts != r.TotalAmount {
		return enums.ErrInvalidAmountCalculated
	}

	return nil
}

type TransactionResponse struct {
	TransactionID   string                `json:"transaction_id"`
	CreatedAt       time.Time             `json:"created_at"`
	UserID          string                `json:"user_id"`
	OrderID         string                `json:"order_id"`
	TransactionType enums.TransactionType `json:"transaction_type"`
	Description     string                `json:"description"`
	TotalAmount     models.Amount         `json:"total_amount"`
	DocumentType    enums.DocumentType    `json:"document_type"`
	Comment         string                `json:"comment"`
	Amount          models.Amount         `json:"amount"`
	Balance         models.Amount         `json:"balance"`
	//Transfers       []TransferResponse    `json:"transfers"`
}

type TransferResponse struct {
	DocumentType enums.DocumentType `json:"document_type"`
	Comment      string             `json:"comment"`
	Amount       models.Amount      `json:"amount"`
	Balance      models.Amount      `json:"balance"`
}

//
//type TransferRequest struct {
//	UserID          string                `json:"user_id"`
//	Type            enums.TransactionType `json:"type"`
//	OrderID         string                `json:"order_id"`
//	SkipLock        bool                  `json:"skip_lock"`
//	DebitAccountID  string                `json:"debit_account_id"`
//	CreditAccountID string                `json:"credit_account_id"`
//	FeeAccountID    string                `json:"fee_account_id"`
//	RawAmount       models.Amount         `json:"raw_amount"`
//	FeeAmount       models.Amount         `json:"fee_amount"`
//	TotalAmount     models.Amount         `json:"total_amount"`
//	Description     string                `json:"description"`
//	FeeDescription  string                `json:"fee_description"`
//}
//
//type TransferResponse struct {
//	CreatedAt       time.Time             `json:"created_at"`
//	TransactionID   string                `json:"transaction_id"`
//	OrderID         string                `json:"order_id"`
//	TransactionType enums.TransactionType `json:"transaction_type"`
//	DocumentType    enums.DocumentType    `json:"document_type"`
//	Description     string                `json:"description"`
//	Comment         string                `json:"comment"`
//	Amount          models.Amount         `json:"amount"`
//	Balance         models.Amount         `json:"balance"`
//}

type ReverseRequest struct {
	TransactionID string `json:"transaction_id"`
	UserID        string `json:"user_id"`
	Description   string `json:"description"`
}

type InquiryRequest struct {
	TransactionID string `json:"transaction_id" query:"transaction_id"`
	UserID        string `json:"user_id" query:"user_id"`
	OrderID       string `json:"order_id" query:"order_id"`
}

type HistoryRequest struct {
	models.PaginationRequest
	UserID    string `json:"user_id"`
	AccountID string `json:"account_id"`
}

type HistoryResponse struct {
	Transactions []*TransactionResponse     `json:"transactions"`
	Meta         *models.PaginationResponse `json:"meta"`
}
