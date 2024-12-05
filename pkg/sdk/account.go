package sdk

import (
	"errors"
	"fmt"
	"github.com/sinameshkini/fingo/internal/models"
	"github.com/sinameshkini/microkit/pkg/utils"
)

func (c *Client) GetAccount(id string) (resp *models.AccountResponse, err error) {
	apiResp := &models.Response{}

	r, err := c.rc.R().
		SetResult(apiResp).
		SetError(apiResp).
		Get(fmt.Sprintf("/accounts/%s", id))
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

func (c *Client) GetAccounts(userID string) (resp []*models.AccountResponse, err error) {
	apiResp := &models.Response{}

	r, err := c.rc.R().
		SetResult(apiResp).
		SetError(apiResp).
		Get(fmt.Sprintf("/accounts?user_id=%s", userID))
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

func (c *Client) CreateAccount(req models.CreateAccount) (resp *models.AccountResponse, err error) {
	apiResp := &models.Response{}

	r, err := c.rc.R().
		SetBody(&req).
		SetResult(apiResp).
		SetError(apiResp).
		Post("/accounts")
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

func (c *Client) GetPolicies(req models.GetSettingsRequest) (resp *models.Settings, err error) {
	apiResp := &models.Response{}

	r, err := c.rc.R().
		SetResult(apiResp).
		SetError(apiResp).
		Get("/policies")
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
