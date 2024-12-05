package main

import (
	"github.com/sinameshkini/fingo/internal/config"
	"github.com/sinameshkini/fingo/pkg/migration"
	"github.com/sinameshkini/microkit/models"
	"github.com/sinameshkini/microkit/pkg/clients/database"
	"log"
)

func main() {
	if err := migrator(); err != nil {
		log.Fatal(err)
	}
}

func migrator() (err error) {
	log.Println("fingo migrator starting ...")

	models.InitSnowflakeID(1)

	db, err := database.NewDBWithConf(config.DefaultConf.Database)
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

	return nil
}
