package sdk

import (
	"errors"
	"fmt"
	"github.com/sinameshkini/fingo/internal/repository/entities"
	"github.com/sinameshkini/fingo/pkg/endpoint"
	"github.com/sinameshkini/microkit/pkg/utils"
)

func (c *Client) GetAccount(id string) (resp *endpoint.AccountResponse, err error) {
	apiResp := &entities.Response{}

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

func (c *Client) GetAccounts(userID string) (resp []*endpoint.AccountResponse, err error) {
	apiResp := &entities.Response{}

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

func (c *Client) CreateAccount(req endpoint.CreateAccount) (resp *endpoint.AccountResponse, err error) {
	apiResp := &entities.Response{}

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

func (c *Client) GetPolicies(req entities.GetSettingsRequest) (resp *entities.Settings, err error) {
	apiResp := &entities.Response{}

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

func (c *Client) GetAccountTypes() (resp []*entities.AccountType, err error) {
	apiResp := &entities.Response{}

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

func (c *Client) GetCurrencies() (resp []*entities.Currency, err error) {
	apiResp := &entities.Response{}

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
