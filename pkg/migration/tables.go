package migration

import (
	"github.com/sinameshkini/fingo/internal/repository/entities"
)

var Tables = []interface{}{
	&entities.Currency{},
	&entities.Account{},
	&entities.AccountType{},
	&entities.Transaction{},
	&entities.Document{},
	&entities.Policy{},
}
