package config

import (
	"github.com/sinameshkini/microkit/pkg/clients/cache"
	"github.com/sinameshkini/microkit/pkg/clients/database"
)

type Config struct {
	Env      *Env
	Address  string
	Database *database.Config
	Cache    *cache.Config
}

type Env struct {
	Lock bool
}

var DefaultConf = Config{
	Env: &Env{
		Lock: true,
	},
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
