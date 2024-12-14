package end2end

import (
	"github.com/sinameshkini/fingo/pkg/enums"
	"github.com/sinameshkini/fingo/test"
	"github.com/sinameshkini/microkit/pkg/utils"
	"testing"
)

// ID:			TS002_Account
// Scenario:	Create account and get account (info and balance)

func Test_Account(t *testing.T) {
	cli, _, err := test.Setup()
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	account, err := CreateAccount(cli, "user1", enums.ACCOUNTTYPEWALLET, "user1")
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	utils.PrintJson(account)
}

func Test_Account_Duplicate(t *testing.T) {
	cli, _, err := test.Setup()
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	account1, err := CreateAccount(cli, "user1", enums.ACCOUNTTYPEWALLET, "user1")
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	utils.PrintJson(account1)

	_, err = CreateAccount(cli, "user1", enums.ACCOUNTTYPEWALLET, "user1")
	if err == nil {
		t.Error("account inserted out of limits")
		t.FailNow()
	}
	utils.PrintJson(err)
}
