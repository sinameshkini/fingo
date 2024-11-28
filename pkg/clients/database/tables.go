package database

import "github.com/sinameshkini/fingo/internal/models"

var tables = []interface{}{
	&models.Currency{},
	&models.Account{},
	&models.AccountType{},
	&models.Transaction{},
	&models.Document{},
	&models.Policy{},
}
