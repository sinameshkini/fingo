package sdk

import (
	"errors"
	"fmt"
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

func (c *Client) Inquiry(req endpoint.InquiryRequest) (resp []*endpoint.TransferResponse, err error) {
	apiResp := &entities.Response{}

	r, err := c.rc.R().
		SetResult(apiResp).
		SetError(apiResp).
		Get(fmt.Sprintf("/inquiry?user_id=%s&transaction_id=%s&order_id=%s", req.UserID, req.TransactionID, req.OrderID))
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

func (c *Client) History(req endpoint.HistoryRequest) (resp *endpoint.HistoryResponse, err error) {
	apiResp := &entities.Response{}

	r, err := c.rc.R().
		SetBody(&req).
		SetResult(apiResp).
		SetError(apiResp).
		Post("/history")
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
