package service

import (
	"github.com/sinameshkini/fingo/pkg/clients/cache"
	"github.com/sinameshkini/fingo/pkg/clients/database"
)

type Config struct {
	Address  string
	Database database.Config
	Cache    cache.Config
}

var DefaultConf = Config{
	Address:  ":4000",
	Database: database.Config{},
	Cache: cache.Config{
		Host: "localhost:6379",
		DB:   0,
	},
}
