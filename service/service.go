package service

import (
	"github.com/sinameshkini/fingo/internal/controller"
	"github.com/sinameshkini/fingo/internal/core"
	"github.com/sinameshkini/fingo/internal/models"
	"github.com/sinameshkini/fingo/internal/repository"
	"github.com/sinameshkini/fingo/pkg/clients/database"
	"log"
)

func Run() (err error) {
	log.Println("fingo starting ...")

	models.InitID(1)

	db, err := database.New(database.Config{})
	if err != nil {
		return
	}

	repo := repository.New(db)

	c := core.New(repo)

	return controller.Init(c)
}
