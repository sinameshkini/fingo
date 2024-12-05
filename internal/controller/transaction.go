package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/sinameshkini/fingo/pkg/endpoint"
)

func (h *ctrl) transfer(c echo.Context) error {
	req := endpoint.TransferRequest{}
	if err := c.Bind(&req); err != nil {
		return responseError(c, err)
	}

	resp, err := h.core.Transfer(c.Request().Context(), req)
	if err != nil {
		return responseError(c, err)
	}

	return response(c, resp)
}

func (h *ctrl) reverse(c echo.Context) error {
	req := endpoint.ReverseRequest{}
	if err := c.Bind(&req); err != nil {
		return responseError(c, err)
	}

	resp, err := h.core.Reverse(c.Request().Context(), req)
	if err != nil {
		return responseError(c, err)
	}

	return response(c, resp)
}

func (h *ctrl) inquiry(c echo.Context) error {
	req := endpoint.InquiryRequest{}
	if err := c.Bind(&req); err != nil {
		return responseError(c, err)
	}

	resp, err := h.core.Inquiry(c.Request().Context(), req)
	if err != nil {
		return responseError(c, err)
	}

	return response(c, resp)
}

func (h *ctrl) history(c echo.Context) error {
	req := endpoint.HistoryRequest{}
	if err := c.Bind(&req); err != nil {
		return responseError(c, err)
	}

	resp, err := h.core.GetTransactions(c.Request().Context(), req)
	if err != nil {
		return responseError(c, err)
	}

	return response(c, resp)
}
