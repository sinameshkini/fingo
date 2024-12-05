package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/sinameshkini/fingo/internal/repository/entities"
)

func (h *ctrl) getPolicy(c echo.Context) error {
	req := entities.GetSettingsRequest{}
	if err := c.Bind(&req); err != nil {
		return responseError(c, err)
	}

	resp, err := h.core.GetSettings(c.Request().Context(), req)
	if err != nil {
		return responseError(c, err)
	}

	return response(c, resp)
}
