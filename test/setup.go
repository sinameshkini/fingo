package test

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/sinameshkini/fingo/internal/config"
	"github.com/sinameshkini/fingo/pkg/migration"
	"github.com/sinameshkini/fingo/pkg/sdk"
	"github.com/sinameshkini/fingo/service"
	"github.com/sinameshkini/microkit/models"
	"github.com/sinameshkini/microkit/pkg/clients/database"
	"time"
)

var conf = config.DefaultConf

func Setup() (cli *sdk.Client, baseURL string, err error) {
	models.InitSnowflakeID(1)

	baseURL = fmt.Sprintf("http://localhost%s/v1", conf.Address)

	db, err := database.NewDBWithConf(conf.Database)
	if err != nil {
		return
	}

	if err = database.Drop(db, migration.Tables); err != nil {
		return
	}

	if err = database.Migrate(db, migration.Tables); err != nil {
		return
	}

	if err = migration.Seed(db); err != nil {
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
