package end2end

import (
	"github.com/sinameshkini/fingo/pkg/endpoint"
	"github.com/sinameshkini/fingo/pkg/enums"
	"github.com/sinameshkini/fingo/test"
	"github.com/sinameshkini/microkit/pkg/utils"
	"testing"
)

// ID:			TS004_Reverse
// Scenario:	Reverse transaction

func Test_Reverse(t *testing.T) {
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

	depositTxn, err := cli.Transfer(endpoint.TransactionRequest{
		UserID:      "admin",
		Type:        enums.Deposit,
		OrderID:     "1234",
		TotalAmount: 1000,
		Description: "Deposit test",
		Transfers: []endpoint.TransferRequest{
			{
				DebitAccountID:  shadowAccount.ID,
				CreditAccountID: account1.ID,
				Amount:          1000,
				Comment:         "Deposit",
			},
		},
	})
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	utils.PrintJson(depositTxn)

	reverseTxn, err := cli.Reverse(endpoint.ReverseRequest{
		UserID:        "admin",
		TransactionID: depositTxn.TransactionID,
	})
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	utils.PrintJson(reverseTxn)

	account1, err = cli.GetAccount(account1.ID)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	if account1.Balance != 0 {
		t.Error("account1 balance should be 0")
	}

	shadowAccount, err = cli.GetAccount(shadowAccount.ID)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	if shadowAccount.Balance != 0 {
		t.Error("shadowAccount balance should be 0")
	}
}
