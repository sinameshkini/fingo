package end2end

import (
	"github.com/sinameshkini/fingo/pkg/enums"
	"github.com/sinameshkini/fingo/test"
	"github.com/sinameshkini/microkit/models"
	"log"
	"testing"
)

// ID:			TS006_History
// Scenario:	Check history response

func Test_History(t *testing.T) {
	var (
		depositCount               = 100
		amount       models.Amount = 10
	)

	cli, baseURL, err := test.Setup()
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	shadowAccount, err := GetAccount(cli, "admin", enums.ACCOUNTTYPESHADOW)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	beforeBalance := shadowAccount.Balance

	if err = NormalActor(baseURL, "user1", shadowAccount.ID, depositCount, amount); err != nil {
		log.Fatalln(err)
		return
	}

	// Check shadow account balance
	shadowAccount, err = cli.GetAccount(shadowAccount.ID)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	if wantBalance := int(beforeBalance) - 1*depositCount*int(amount); int(shadowAccount.Balance) != wantBalance {
		t.Errorf("want %d , got %d", wantBalance, shadowAccount.Balance)
		t.FailNow()
	}
}
