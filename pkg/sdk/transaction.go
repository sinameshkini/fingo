package sdk

import (
	"errors"
	"github.com/sinameshkini/fingo/internal/repository/entities"
	"github.com/sinameshkini/fingo/pkg/endpoint"
	"github.com/sinameshkini/microkit/pkg/utils"
)

func (c *Client) Transfer(req endpoint.TransferRequest) (resp *endpoint.TransferResponse, err error) {
	apiResp := &entities.Response{}

	r, err := c.rc.R().
		SetBody(&req).
		SetResult(apiResp).
		SetError(apiResp).
		Post("/transfer")
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

func (c *Client) Reverse(req endpoint.ReverseRequest) (resp *endpoint.TransferResponse, err error) {
	apiResp := &entities.Response{}

	r, err := c.rc.R().
		SetBody(&req).
		SetResult(apiResp).
		SetError(apiResp).
		Post("/reverse")
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
