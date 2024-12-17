package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/sinameshkini/fingo/internal/config"
	"github.com/sinameshkini/fingo/internal/core"
	"github.com/sinameshkini/fingo/internal/repository/entities"
	"github.com/sinameshkini/fingo/pkg/enums"
	"net/http"
)

type ctrl struct {
	core *core.Core
}

func Init(conf config.Config, c *core.Core) error {
	e := echo.New()
	h := &ctrl{core: c}

	api := e.Group("/v1")

	api.GET("/status", func(c echo.Context) error {
		return c.String(http.StatusOK, "Running")
	})

	// Configuration routes
	api.GET("/account_types", h.accountTypes)
	api.GET("/currencies", h.currencies)
	api.GET("/settings", h.getSettings)

	// Account routes
	api.POST("/accounts", h.newAccount)
	api.GET("/accounts/:id", h.getAccount)
	api.GET("/accounts", h.getAccounts)

	// Transaction routes
	api.POST("/transfer", h.transfer)
	api.POST("/reverse", h.reverse)
	api.GET("/inquiry", h.inquiry)
	api.POST("/history", h.history)

	return e.Start(conf.Address)
}

func response(c echo.Context, payload any) error {
	return c.JSON(http.StatusOK, entities.Response{
		Code:    0,
		Message: "success",
		Data:    payload,
	})
}

func accept(c echo.Context, payload any) error {
	return c.JSON(http.StatusAccepted, entities.Response{
		Code:    0,
		Message: "accepted",
		Data:    payload,
	})
}

func responseError(c echo.Context, err error) error {
	resp := entities.Response{
		Code:    1,
		Message: err.Error(),
	}

	httpCode, ok := enums.ErrHTTPCode[err]
	if !ok {
		httpCode = http.StatusInternalServerError
	}

	if errCode, ok := enums.ErrCode[err]; ok {
		resp.Code = errCode
	}

	return c.JSON(httpCode, resp)
}
