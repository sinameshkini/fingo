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
	var (
		count uint
	)

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

	accounts, err := c.repo.GetAccounts(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	for _, a := range accounts {
		if a.AccountTypeID == req.AccountTypeID {
			count++
		}
	}

	settings, err := c.GetSettings(ctx, models.GetSettingsRequest{
		UserID:        req.UserID,
		AccountTypeID: req.AccountTypeID,
	})
	if err != nil {
		return
	}

	if count >= settings.Limits.NumberOfAccounts[req.AccountTypeID] {
		return nil, models.ErrPermissionDenied
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

	return c.GetAccount(ctx, account.ID)
}

func (c *Core) GetAccount(ctx context.Context, id models.ID) (resp *models.AccountResponse, err error) {
	account, err := c.repo.GetAccount(ctx, id)
	if err != nil {
		return
	}

	balance, err := c.repo.GetBalance(ctx, id)

	resp = account.ToResponse(balance)
	return
}

func (c *Core) GetAccounts(ctx context.Context, userID string) (resp []*models.AccountResponse, err error) {
	accounts, err := c.repo.GetAccounts(ctx, userID)
	if err != nil {
		return
	}

	for _, a := range accounts {
		balance, _ := c.repo.GetBalance(ctx, a.ID)
		resp = append(resp, a.ToResponse(balance))
	}

	return
}
