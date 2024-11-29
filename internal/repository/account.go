package repository

import (
	"context"
	"github.com/sinameshkini/fingo/internal/models"
)

func (r *repo) CreateAccount(ctx context.Context, account models.Account) (resp *models.Account, err error) {
	if err = r.db.WithContext(ctx).Create(&account).Error; err != nil {
		return
	}

	return &account, nil
}

func (r *repo) GetAccount(ctx context.Context, id models.ID) (resp *models.Account, err error) {
	if err = r.db.WithContext(ctx).
		Preload("AccountType").
		Preload("Currency").
		First(&resp, id).Error; err != nil {
		return
	}

	return
}

func (r *repo) GetAccounts(ctx context.Context, userID string) (resp []*models.Account, err error) {
	if err = r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Preload("AccountType").
		Preload("Currency").
		Find(&resp).Error; err != nil {
		return
	}

	return
}

func (r *repo) GetAccountTypes(ctx context.Context) (resp []*models.AccountType, err error) {
	if err = r.db.WithContext(ctx).Find(&resp).Error; err != nil {
		return
	}

	return
}

func (r *repo) GetAccountType(ctx context.Context, id string) (resp *models.AccountType, err error) {
	if err = r.db.WithContext(ctx).First(&resp, id).Error; err != nil {
		return
	}

	return
}

func (r *repo) GetCurrencies(ctx context.Context) (resp []*models.Currency, err error) {
	if err = r.db.WithContext(ctx).Find(&resp).Error; err != nil {
		return
	}

	return
}

func (r *repo) GetCurrency(ctx context.Context, id uint) (resp *models.Currency, err error) {
	if err = r.db.WithContext(ctx).First(&resp, id).Error; err != nil {
		return
	}

	return
}
