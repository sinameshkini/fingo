package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/sinameshkini/fingo/pkg/endpoint"
)

func (h *ctrl) getSettings(c echo.Context) error {
	req := endpoint.GetSettingsRequest{}
	if err := c.Bind(&req); err != nil {
		return responseError(c, err)
	}

	resp, err := h.core.GetSettings(c.Request().Context(), req)
	if err != nil {
		return responseError(c, err)
	}

	return response(c, resp)
}
