package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/sinameshkini/fingo/internal/config"
	"github.com/sinameshkini/fingo/internal/core"
	"github.com/sinameshkini/fingo/internal/models"
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

	// Account routes
	api.GET("/account_types", h.accountTypes)
	api.GET("/currencies", h.currencies)

	api.POST("/accounts", h.newAccount)
	api.GET("/accounts/:id", h.getAccount)
	api.GET("/accounts", h.getAccounts)

	api.GET("/policies", h.getPolicy)

	api.POST("/transfer", h.transfer)
	api.POST("/reverse", h.reverse)

	return e.Start(conf.Address)
}

func response(c echo.Context, payload any) error {
	return c.JSON(http.StatusOK, models.Response{
		Code:    0,
		Message: "success",
		Data:    payload,
	})
}

func responseError(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, models.Response{
		Code:    1,
		Message: err.Error(),
		Data:    nil,
	})
}
