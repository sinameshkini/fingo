package core

import (
	"context"
	"github.com/sinameshkini/fingo/internal/models"
	"log"
)

func (c *Core) GetAccountTypes(ctx context.Context) ([]*models.AccountType, error) {
	return c.repo.GetAccountTypes(ctx)
}

func (c *Core) GetCurrencies(ctx context.Context) ([]*models.Currency, error) {
	return c.repo.GetCurrencies(ctx)
}

func (c *Core) NewAccount(ctx context.Context, req models.CreateAccount) (resp *models.AccountResponse, err error) {
	_, err = c.repo.GetAccountType(ctx, req.AccountTypeID)
	if err != nil {
		log.Println(err)
		return nil, models.ErrAccountTypeInvalid
	}

	_, err = c.repo.GetCurrency(ctx, req.CurrencyID)
	if err != nil {
		log.Println(err)
		return nil, models.ErrCurrencyInvalid
	}

	account := &models.Account{
		UserID:        req.UserID,
		AccountTypeID: req.AccountTypeID,
		CurrencyID:    req.CurrencyID,
		Name:          req.Name,
		IsDefault:     false,
		IsEnable:      true,
	}

	account, err = c.repo.CreateAccount(ctx, *account)
	if err != nil {
		log.Println(err)
		return nil, models.ErrInternal
	}

	account, err = c.repo.GetAccount(ctx, account.ID)
	if err != nil {
		log.Println(err)
		return nil, models.ErrInternal
	}

	return account.ToResponse(), nil
}
