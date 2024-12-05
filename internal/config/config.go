package config

import (
	"github.com/sinameshkini/microkit/pkg/clients/cache"
	"github.com/sinameshkini/microkit/pkg/clients/database"
)

type Config struct {
	Address  string
	Database *database.Config
	Cache    *cache.Config
}

var DefaultConf = Config{
	Address: ":4000",
	Database: &database.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "admin",
		Password: "admin",
		Schema:   "fingo",
		Debug:    true,
	},
	Cache: &cache.Config{
		Host: "localhost:6379",
		DB:   0,
	},
}
