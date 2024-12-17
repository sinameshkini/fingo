package repository

import (
	"context"
	"github.com/sinameshkini/fingo/internal/repository/entities"
	"github.com/sinameshkini/fingo/pkg/endpoint"
	"github.com/sinameshkini/fingo/pkg/enums"
	"github.com/sinameshkini/microkit/models"
	"gorm.io/gorm"
	"sync"
)

var (
	r    *repo
	once sync.Once
)

type repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	once.Do(func() {
		r = &repo{
			db: db,
		}
	})

	return r
}

type Repository interface {
	CreateAccount(ctx context.Context, account entities.Account) (*entities.Account, error)
	GetAccount(ctx context.Context, id models.SID) (*entities.Account, error)
	GetAccounts(ctx context.Context, userID string) ([]*entities.Account, error)
	GetBalance(ctx context.Context, accountID models.SID) (models.Amount, error)

	NewTransaction(ctx context.Context, userID, orderID, description string, txnType enums.TransactionType, amount models.Amount) (txn *entities.Transaction, tx *gorm.DB, err error)
	CommitTransaction(tx *gorm.DB) error
	Transfer(tx *gorm.DB, amount models.Amount, txnID, debID, credID models.SID, comment string) error
	Reverse(tx *gorm.DB, transaction *entities.Transaction, reverseTxnID models.SID) error

	GetTransaction(ctx context.Context, txnID models.SID) (*entities.Transaction, error)
	Inquiry(ctx context.Context, req endpoint.InquiryRequest) (resp []*entities.Transaction, err error)
	GetByOrderID(ctx context.Context, userID, orderID string) (resp []*entities.Transaction, err error)
	GetHistory(ctx context.Context, req endpoint.HistoryRequest) (resp []*entities.Document, meta *models.PaginationResponse, err error)

	FetchPolicies(ctx context.Context, req endpoint.FetchPoliciesRequest) (resp []*entities.Policy, meta *models.PaginationResponse, err error)
	CreatePolicy(ctx context.Context, req entities.Policy) (resp *entities.Policy, err error)
	UpdatePolicy(ctx context.Context, policyID models.SID, req entities.Policy) (resp *entities.Policy, err error)
	DeletePolicy(ctx context.Context, policyID models.SID) (err error)

	GetPolicies(ctx context.Context, userID, accountID, accountType string) ([]*entities.Policy, error)
	GetAccountTypes(ctx context.Context) ([]*entities.AccountType, error)
	GetAccountType(ctx context.Context, id string) (*entities.AccountType, error)
	GetCurrencies(ctx context.Context) ([]*entities.Currency, error)
	GetCurrency(ctx context.Context, id uint) (*entities.Currency, error)
}
