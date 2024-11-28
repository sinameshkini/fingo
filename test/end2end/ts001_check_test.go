package end2end

import (
	"github.com/sinameshkini/fingo/pkg/utils"
	"github.com/sinameshkini/fingo/test"
	"testing"
)

// ID:			TS001_Check
// Scenario:	Check all required system settings

func Test_TS001_Check(t *testing.T) {
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
	utils.PrintJson(currencies)

	accountTypes, err := cli.GetAccountTypes()
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	utils.PrintJson(accountTypes)
}
