package end2end

import (
	"github.com/sinameshkini/fingo/pkg/endpoint"
	"github.com/sinameshkini/fingo/pkg/enums"
	"github.com/sinameshkini/fingo/test"
	"github.com/sinameshkini/microkit/pkg/utils"
	"testing"
)

// ID:			TS006_Transfer_With_Fee
// Scenario:	Deposit transfer with fee

func Test_TS006_Transfer_With_Fee(t *testing.T) {
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

	shadowAccount, err := GetAccount(cli, "admin", enums.ACCOUNTTYPESHADOW)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	feeAccount, err := GetAccount(cli, "admin", enums.ACCOUNTTYPEFEE)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	depositTxn, err := cli.Transfer(endpoint.TransferRequest{
		UserID:          "admin",
		Type:            enums.Deposit,
		OrderID:         "1234",
		DebitAccountID:  shadowAccount.ID,
		CreditAccountID: account1.ID,
		RawAmount:       1000,
		Description:     "Deposit test",
		FeeAccountID:    feeAccount.ID,
		FeeAmount:       100,
		FeeDescription:  "Deposit Fee",
		TotalAmount:     1100,
	})
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	utils.PrintJson(depositTxn)

	account1, err = cli.GetAccount(account1.ID)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	if account1.Balance != 1000 {
		t.Error("account1 balance should be 1000")
	}

	shadowAccount, err = cli.GetAccount(shadowAccount.ID)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	if shadowAccount.Balance != -1100 {
		t.Error("shadowAccount balance should be -1100")
	}

	feeAccount, err = cli.GetAccount(feeAccount.ID)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	if feeAccount.Balance != 100 {
		t.Error("feeAccount balance should be 100")
	}
}
