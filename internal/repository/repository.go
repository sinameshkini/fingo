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
	GetAccountTypes(ctx context.Context) ([]*models.AccountType, error)
	GetAccountType(ctx context.Context, id string) (*models.AccountType, error)
	GetCurrencies(ctx context.Context) ([]*models.Currency, error)
	GetCurrency(ctx context.Context, id uint) (*models.Currency, error)

	CreateAccount(ctx context.Context, account models.Account) (*models.Account, error)
	GetAccount(ctx context.Context, id models.ID) (*models.Account, error)
}
