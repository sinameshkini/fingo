package end2end

import (
	"github.com/sinameshkini/fingo/pkg/endpoint"
	"github.com/sinameshkini/fingo/test"
	"github.com/sinameshkini/microkit/pkg/utils"
	"testing"
)

// ID:			TS001_Check
// Scenario:	Check all required system settings

func Test_TS001_Check(t *testing.T) {
	cli, _, err := test.Setup()
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	settings, err := cli.GetSettings(endpoint.GetSettingsRequest{})
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	utils.PrintJson(settings)

	currencies, err := cli.GetCurrencies()
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	utils.PrintJson(currencies)

	accountTypes, err := cli.GetAccountTypes()
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	utils.PrintJson(accountTypes)

	accounts, err := cli.GetAccounts("admin")
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	utils.PrintJson(accounts)
}
