package end2end

import (
	"errors"
	"github.com/sinameshkini/fingo/internal/models"
	"github.com/sinameshkini/fingo/pkg/sdk"
)

func createAccount(cli *sdk.Client, userID, accountType, name string) (resp *models.AccountResponse, err error) {
	currencies, err := cli.GetCurrencies()
	if err != nil {
		return
	}

	accountTypes, err := cli.GetAccountTypes()
	if err != nil {
		return
	}

	if len(currencies) == 0 || len(accountTypes) == 0 {
		return nil, errors.New("no currencies or accountTypes")
	}

	account, err := cli.CreateAccount(models.CreateAccount{
		UserID:        userID,
		AccountTypeID: accountType,
		CurrencyID:    currencies[0].ID,
		Name:          name,
	})
	if err != nil {
		return
	}

	return cli.GetAccount(account.ID)
}

func getAccount(cli *sdk.Client, userID, accountType string) (resp *models.AccountResponse, err error) {
	adminAccounts, err := cli.GetAccounts(userID)
	if err != nil {
		return
	}

	for _, aa := range adminAccounts {
		if aa.AccountType.ID == accountType {
			return aa, nil
		}
	}

	return nil, models.ErrNotFound
}
