package repository

import (
	"context"
	"github.com/sinameshkini/fingo/internal/models"
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
	CreateAccount(ctx context.Context, account models.Account) (*models.Account, error)
	GetAccount(ctx context.Context, id models.ID) (*models.Account, error)
	GetAccounts(ctx context.Context, userID string) ([]*models.Account, error)
	GetBalance(ctx context.Context, accountID models.ID) (models.Amount, error)

	NewTransaction(ctx context.Context) *gorm.DB
	CommitTransaction(tx *gorm.DB) error
	Initial(*gorm.DB, string, string, models.TransactionType, models.Amount, string) (*models.Transaction, error)
	Transfer(tx *gorm.DB, amount models.Amount, txnID, debID, credID models.ID, comment string) error
	Reverse(tx *gorm.DB, transaction *models.Transaction, reverseTxnID models.ID) error

	GetTransaction(ctx context.Context, txnID models.ID) (*models.Transaction, error)
	GetHistory(ctx context.Context, accountID models.ID) ([]*models.Document, error)

	GetPolicies(ctx context.Context, userID, accountID, accountType string) ([]*models.Policy, error)
	GetAccountTypes(ctx context.Context) ([]*models.AccountType, error)
	GetAccountType(ctx context.Context, id string) (*models.AccountType, error)
	GetCurrencies(ctx context.Context) ([]*models.Currency, error)
	GetCurrency(ctx context.Context, id uint) (*models.Currency, error)
}
