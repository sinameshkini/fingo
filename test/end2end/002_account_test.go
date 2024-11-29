package end2end

import (
	"github.com/sinameshkini/fingo/internal/models"
	"github.com/sinameshkini/fingo/pkg/utils"
	"github.com/sinameshkini/fingo/test"
	"testing"
)

// ID:			TS002_Account
// Scenario:	Create account and get account (info and balance)

func Test_TS002_Account(t *testing.T) {
	cli, err := test.Setup()
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	currencies, err := cli.GetCurrencies()
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	accountTypes, err := cli.GetAccountTypes()
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	if len(currencies) == 0 || len(accountTypes) == 0 {
		t.Error("currency and account types are empty")
		t.FailNow()
	}

	account, err := cli.CreateAccount(models.CreateAccount{
		UserID:        "1234",
		AccountTypeID: accountTypes[1].ID,
		CurrencyID:    currencies[0].ID,
		Name:          "Test",
	})
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	account, err = cli.GetAccount(account.ID)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	utils.PrintJson(account)
}
