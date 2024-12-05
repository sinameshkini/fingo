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
	"time"
)

func Run(conf config.Config) (err error) {
	log.Println("fingo starting ...")

	models.InitSnowflakeID(1)

	db, err := database.NewDBWithConf(conf.Database)
	if err != nil {
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		return
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(80)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Println(err)
		}
	}()

	ca := cache.New(*conf.Cache)

	repo := repository.New(db)

	c := core.New(repo, ca)

	return controller.Init(conf, c)
}
