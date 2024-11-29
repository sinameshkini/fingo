package end2end

import (
	"github.com/sinameshkini/fingo/internal/models"
	"github.com/sinameshkini/fingo/pkg/utils"
	"github.com/sinameshkini/fingo/test"
	"testing"
)

// ID:			TS004_Reverse
// Scenario:	Reverse transaction

func Test_TS004_Reverse(t *testing.T) {
	cli, err := test.Setup()
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	account1, err := createAccount(cli, "user1", models.ACCOUNTTYPEWALLET, "user1")
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	shadowAccount, err := getAccount(cli, "admin", models.ACCOUNTTYPESHADOW)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	depositTxn, err := cli.Transfer(models.TransferRequest{
		UserID:          "admin",
		Type:            models.Deposit,
		OrderID:         "1234",
		DebitAccountID:  shadowAccount.ID,
		CreditAccountID: account1.ID,
		RawAmount:       1000,
		TotalAmount:     1000,
		Description:     "Deposit test",
	})
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	utils.PrintJson(depositTxn)

	reverseTxn, err := cli.Reverse(models.ReverseRequest{
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
