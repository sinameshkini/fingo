package end2end

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sinameshkini/fingo/internal/models"
	"github.com/sinameshkini/fingo/test"
	"github.com/sinameshkini/microkit/pkg/utils"
	"testing"
)

// ID:			TS005_Load
// Scenario:	High load request (TPS calculate)

func Test_TS005_Load(t *testing.T) {
	var (
		userCount    models.Amount = 30
		depositCount models.Amount = 20
		amount       models.Amount = 10
		accountMap                 = make(map[string]*models.AccountResponse)
	)

	cli, err := test.Setup()
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	shadowAccount, err := getAccount(cli, "admin", models.ACCOUNTTYPESHADOW)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	for i := models.Amount(0); i < userCount; i++ {
		userID := fmt.Sprintf("user%d", i)
		account, err := createAccount(cli, userID, models.ACCOUNTTYPEWALLET, userID)
		if err != nil {
			t.Error(err.Error())
			t.FailNow()
		}

		accountMap[userID] = account

		for i := models.Amount(0); i < depositCount; i++ {
			depositTxn, err := cli.Transfer(models.TransferRequest{
				UserID:          "admin",
				Type:            models.Deposit,
				OrderID:         uuid.NewString(),
				DebitAccountID:  shadowAccount.ID,
				CreditAccountID: account.ID,
				RawAmount:       amount,
				TotalAmount:     amount,
				Description:     fmt.Sprintf("%s deposit %d", userID, i),
			})
			if err != nil {
				t.Error(err.Error())
				t.FailNow()
			}

			utils.PrintJson(depositTxn)
		}
	}

	for _, account := range accountMap {
		account, err = cli.GetAccount(account.ID)
		if err != nil {
			t.Error(err.Error())
			t.FailNow()
		}

		if wantBalance := amount * depositCount; account.Balance != wantBalance {
			t.Error("want", wantBalance, "got", account.Balance)
		}
	}

	shadowAccount, err = cli.GetAccount(shadowAccount.ID)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	if wantBalance := -1 * depositCount * userCount * amount; shadowAccount.Balance != wantBalance {
		t.Error("want", wantBalance, "got", shadowAccount.Balance)
	}
}
