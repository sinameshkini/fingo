package service

import (
	"github.com/sinameshkini/fingo/internal/config"
	"github.com/sinameshkini/fingo/internal/controller"
	"github.com/sinameshkini/fingo/internal/core"
	"github.com/sinameshkini/fingo/internal/repository"
	"github.com/sinameshkini/microkit/models"
	"github.com/sinameshkini/microkit/pkg/clients/cache"
	"github.com/sinameshkini/microkit/pkg/clients/database"
	"log"
)

func Run(conf config.Config) (err error) {
	log.Println("fingo starting ...")

	models.InitSnowflakeID(1)

	db, err := database.NewDBWithConf(conf.Database)
	if err != nil {
		return
	}

	ca := cache.New(*conf.Cache)

	repo := repository.New(db)

	c := core.New(repo, ca)

	return controller.Init(conf, c)
}
