package end2end

import (
	"github.com/sinameshkini/fingo/internal/models"
	"github.com/sinameshkini/fingo/test"
	"github.com/sinameshkini/microkit/pkg/utils"
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

	account, err := createAccount(cli, "user1", models.ACCOUNTTYPEWALLET, "user1")
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	utils.PrintJson(account)
}
