package sdk

import (
	"errors"
	"github.com/sinameshkini/fingo/internal/models"
	"github.com/sinameshkini/microkit/pkg/utils"
)

func (c *Client) Transfer(req models.TransferRequest) (resp *models.TransferResponse, err error) {
	apiResp := &models.Response{}

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

func (c *Client) Reverse(req models.ReverseRequest) (resp *models.TransferResponse, err error) {
	apiResp := &models.Response{}

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
