package sdk

import (
	"errors"
	"github.com/sinameshkini/fingo/internal/models"
	"github.com/sinameshkini/fingo/pkg/utils"
)

func (c *Client) GetAccountTypes() (resp []*models.AccountType, err error) {
	apiResp := &models.Response{}

	r, err := c.rc.R().
		SetResult(apiResp).
		SetError(apiResp).
		Get("/account_types")
	if err != nil {
		return
	}

	if r.IsError() {
		err = errors.New(r.String())
		return
	}

	if err = utils.JsonAssertion(apiResp.Data, &resp); err != nil {
		return
	}

	return
}

func (c *Client) GetCurrencies() (resp []*models.Currency, err error) {
	apiResp := &models.Response{}

	r, err := c.rc.R().
		SetResult(apiResp).
		SetError(apiResp).
		Get("/currencies")
	if err != nil {
		return
	}

	if r.IsError() {
		err = errors.New(r.String())
		return
	}

	if err = utils.JsonAssertion(apiResp.Data, &resp); err != nil {
		return
	}

	return
}
