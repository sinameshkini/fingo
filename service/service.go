package service

import (
	"github.com/sinameshkini/fingo/internal/controller"
	"github.com/sinameshkini/fingo/internal/core"
	"github.com/sinameshkini/fingo/internal/models"
	"github.com/sinameshkini/fingo/internal/repository"
	"github.com/sinameshkini/fingo/pkg/clients/cache"
	"github.com/sinameshkini/fingo/pkg/clients/database"
	"log"
)

func Run(conf Config) (err error) {
	log.Println("fingo starting ...")

	models.InitID(1)

	db, err := database.New(conf.Database)
	if err != nil {
		return
	}

	ca := cache.New(conf.Cache)

	repo := repository.New(db)

	c := core.New(repo, ca)

	return controller.Init(c)
}
