package service

import "github.com/sinameshkini/fingo/pkg/clients/database"

type Config struct {
	Address  string
	Database database.Config
}
