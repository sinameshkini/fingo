package migration

import "github.com/sinameshkini/fingo/internal/models"

var Tables = []interface{}{
	&models.Currency{},
	&models.Account{},
	&models.AccountType{},
	&models.Transaction{},
	&models.Document{},
	&models.Policy{},
}
