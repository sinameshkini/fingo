package end2end

import (
	"github.com/sinameshkini/fingo/pkg/endpoint"
	"github.com/sinameshkini/fingo/pkg/enums"
	"github.com/sinameshkini/fingo/test"
	"github.com/sinameshkini/microkit/pkg/utils"
	"testing"
)

// ID:			TS003_Transfer
// Scenario:	Transfer amount between accounts

func Test_Transfer(t *testing.T) {
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

	account2, err := CreateAccount(cli, "user2", enums.ACCOUNTTYPEWALLET, "user2")
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	transferTxn, err := cli.Transfer(endpoint.TransactionRequest{
		UserID:      "user1",
		Type:        enums.Transfer,
		OrderID:     "1234",
		TotalAmount: 400,
		Description: "Transfer test",
		Transfers: []endpoint.TransferRequest{
			{
				DebitAccountID:  account1.ID,
				CreditAccountID: account2.ID,
				Amount:          400,
				Comment:         "Transfer",
			},
		},
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

func Test_Transfer_With_Fee(t *testing.T) {
	cli, _, err := test.Setup()
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	userID := "user1"

	account1, err := CreateAccount(cli, userID, enums.ACCOUNTTYPEWALLET, userID)
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

	depositTxn, err := cli.Transfer(endpoint.TransactionRequest{
		UserID:      "admin",
		Type:        enums.Deposit,
		OrderID:     "1234",
		TotalAmount: 1100,
		Description: "Deposit test",
		Transfers: []endpoint.TransferRequest{
			{
				DebitAccountID:  shadowAccount.ID,
				CreditAccountID: account1.ID,
				Amount:          1000,
				Comment:         "Deposit",
			},
			{
				DebitAccountID:  shadowAccount.ID,
				CreditAccountID: feeAccount.ID,
				Amount:          100,
				Comment:         "Deposit Fee",
			},
		},
	})
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	inquiryResp, err := cli.Inquiry(endpoint.InquiryRequest{
		UserID:  "admin",
		OrderID: depositTxn.OrderID,
	})
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	utils.PrintJson(inquiryResp)

	historyResp, err := cli.History(endpoint.HistoryRequest{
		UserID:    "admin",
		AccountID: shadowAccount.ID,
	})
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	utils.PrintJson(historyResp)

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
