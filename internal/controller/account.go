package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/sinameshkini/fingo/internal/models"
)

func (h *ctrl) accountTypes(c echo.Context) error {
	resp, err := h.core.GetAccountTypes(c.Request().Context())
	if err != nil {
		return responseError(c, err)
	}

	return response(c, resp)
}

func (h *ctrl) currencies(c echo.Context) error {
	resp, err := h.core.GetCurrencies(c.Request().Context())
	if err != nil {
		return responseError(c, err)
	}

	return response(c, resp)
}

func (h *ctrl) newAccount(c echo.Context) error {
	req := models.CreateAccount{}
	if err := c.Bind(&req); err != nil {
		return responseError(c, err)
	}

	resp, err := h.core.NewAccount(c.Request().Context(), req)
	if err != nil {
		return responseError(c, err)
	}

	return response(c, resp)
}
