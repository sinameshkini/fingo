package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/sinameshkini/fingo/internal/repository/entities"
	"github.com/sinameshkini/fingo/pkg/endpoint"
	"github.com/sinameshkini/microkit/models"
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

	return response(c, resp, nil)
}

func (h *ctrl) fetchPolicies(c echo.Context) error {
	req := endpoint.FetchPoliciesRequest{}
	if err := c.Bind(&req); err != nil {
		return responseError(c, err)
	}

	resp, meta, err := h.core.FetchPolicies(c.Request().Context(), req)
	if err != nil {
		return responseError(c, err)
	}

	return response(c, resp, meta)
}

func (h *ctrl) createPolicy(c echo.Context) error {
	req := entities.Policy{}
	if err := c.Bind(&req); err != nil {
		return responseError(c, err)
	}

	resp, err := h.core.CreatePolicy(c.Request().Context(), req)
	if err != nil {
		return responseError(c, err)
	}

	return response(c, resp, nil)
}

func (h *ctrl) updatePolicy(c echo.Context) error {
	req := entities.Policy{}
	if err := c.Bind(&req); err != nil {
		return responseError(c, err)
	}

	policyID := models.ParseSIDf(c.Param("id"))

	resp, err := h.core.UpdatePolicy(c.Request().Context(), policyID, req)
	if err != nil {
		return responseError(c, err)
	}

	return response(c, resp, nil)
}

func (h *ctrl) deletePolicy(c echo.Context) error {
	policyID := models.ParseSIDf(c.Param("id"))

	err := h.core.DeletePolicy(c.Request().Context(), policyID)
	if err != nil {
		return responseError(c, err)
	}

	return response(c, policyID, nil)
}
