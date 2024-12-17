package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/sinameshkini/fingo/pkg/endpoint"
	"github.com/sinameshkini/microkit/models"
)

func (h *ctrl) accountTypes(c echo.Context) error {
	resp, err := h.core.GetAccountTypes(c.Request().Context())
	if err != nil {
		return responseError(c, err)
	}

	return response(c, resp, nil)
}

func (h *ctrl) currencies(c echo.Context) error {
	resp, err := h.core.GetCurrencies(c.Request().Context())
	if err != nil {
		return responseError(c, err)
	}

	return response(c, resp, nil)
}

func (h *ctrl) newAccount(c echo.Context) error {
	req := endpoint.CreateAccount{}
	if err := c.Bind(&req); err != nil {
		return responseError(c, err)
	}

	resp, err := h.core.NewAccount(c.Request().Context(), req)
	if err != nil {
		return responseError(c, err)
	}

	return response(c, resp, nil)
}

func (h *ctrl) getAccount(c echo.Context) error {
	accountID, err := models.ParseSID(c.Param("id"))
	if err != nil {
		return responseError(c, err)
	}

	resp, err := h.core.GetAccount(c.Request().Context(), accountID)
	if err != nil {
		return responseError(c, err)
	}

	return response(c, resp, nil)
}

func (h *ctrl) getAccounts(c echo.Context) error {
	userID := c.QueryParam("user_id")

	resp, err := h.core.GetAccounts(c.Request().Context(), userID)
	if err != nil {
		return responseError(c, err)
	}

	return response(c, resp, nil)
}
