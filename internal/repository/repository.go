package repository

import (
	"context"
	"github.com/sinameshkini/fingo/internal/repository/entities"
	"github.com/sinameshkini/fingo/pkg/enums"
	"github.com/sinameshkini/fingo/pkg/types"
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
	GetBalance(ctx context.Context, accountID models.SID) (types.Amount, error)

	NewTransaction(ctx context.Context) *gorm.DB
	CommitTransaction(tx *gorm.DB) error
	Initial(*gorm.DB, string, string, enums.TransactionType, types.Amount, string) (*entities.Transaction, error)
	Transfer(tx *gorm.DB, amount types.Amount, txnID, debID, credID models.SID, comment string) error
	Reverse(tx *gorm.DB, transaction *entities.Transaction, reverseTxnID models.SID) error

	GetTransaction(ctx context.Context, txnID models.SID) (*entities.Transaction, error)
	GetHistory(ctx context.Context, accountID models.SID) ([]*entities.Document, error)

	GetPolicies(ctx context.Context, userID, accountID, accountType string) ([]*entities.Policy, error)
	GetAccountTypes(ctx context.Context) ([]*entities.AccountType, error)
	GetAccountType(ctx context.Context, id string) (*entities.AccountType, error)
	GetCurrencies(ctx context.Context) ([]*entities.Currency, error)
	GetCurrency(ctx context.Context, id uint) (*entities.Currency, error)
}
