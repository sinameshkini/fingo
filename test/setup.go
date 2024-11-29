package test

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/sinameshkini/fingo/internal/models"
	"github.com/sinameshkini/fingo/pkg/clients/cache"
	"github.com/sinameshkini/fingo/pkg/clients/database"
	"github.com/sinameshkini/fingo/pkg/sdk"
	"github.com/sinameshkini/fingo/service"
	"time"
)

var (
	conf = service.Config{
		Address:  ":4000",
		Database: database.Config{},
		Cache: cache.Config{
			Host: "localhost:6379",
			DB:   0,
		},
	}
)

func Setup() (cli *sdk.Client, err error) {
	models.InitID(1)

	baseURL := fmt.Sprintf("http://localhost%s/v1", conf.Address)

	db, err := database.New(conf.Database)
	if err != nil {
		return
	}

	if err = database.Drop(db); err != nil {
		return
	}

	if err = database.Migrate(db); err != nil {
		return
	}

	if err = database.Seed(db); err != nil {
		return
	}

	go func() {
		if err := service.Run(conf); err != nil {
			return
		}
	}()

	for i := 0; i < 60; i++ {
		resp, err := resty.New().SetDebug(true).R().Get(baseURL + "/status")
		if err == nil && resp.IsSuccess() {
			break
		}

		time.Sleep(time.Second)
	}

	cli = sdk.New(baseURL)

	return
}
