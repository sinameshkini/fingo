package end2end

import (
	"github.com/sinameshkini/fingo/internal/models"
	"github.com/sinameshkini/fingo/pkg/utils"
	"github.com/sinameshkini/fingo/test"
	"testing"
)

// ID:			TS003_Transfer
// Scenario:	Transfer amount between accounts

func Test_TS003_Transfer(t *testing.T) {
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

	account2, err := createAccount(cli, "user2", models.ACCOUNTTYPEWALLET, "user2")
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	transferTxn, err := cli.Transfer(models.TransferRequest{
		UserID:          "user1",
		Type:            models.Transfer,
		OrderID:         "1234",
		DebitAccountID:  account1.ID,
		CreditAccountID: account2.ID,
		RawAmount:       400,
		TotalAmount:     400,
		Description:     "Transfer test",
	})
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	utils.PrintJson(transferTxn)

	account1, err = cli.GetAccount(account1.ID)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	if account1.Balance != 600 {
		t.Error("account1 balance should be 600")
	}

	account2, err = cli.GetAccount(account2.ID)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	if account2.Balance != 400 {
		t.Error("account2 balance should be 400")
	}

	shadowAccount, err = cli.GetAccount(shadowAccount.ID)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	if shadowAccount.Balance != -1000 {
		t.Error("shadowAccount balance should be -1000")
	}
}
