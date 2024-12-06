package end2end

import (
	"github.com/sinameshkini/fingo/pkg/enums"
	"github.com/sinameshkini/fingo/test"
	"github.com/sinameshkini/microkit/models"
	"log"
	"sync"
	"testing"
	"time"
)

// ID:			TS005_Load
// Scenario:	High load request (TPS calculate)

func Test_Load(t *testing.T) {
	var (
		userCount                  = 100
		depositCount               = 5
		amount       models.Amount = 10
		wg           sync.WaitGroup
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

	// Concurrent account creation
	for i := 0; i < userCount; i++ {
		wg.Add(1)
		userID := MakeUserID(i)
		go func() {
			time.Sleep(100 * time.Millisecond)
			defer wg.Done()
			if err = NormalActor(baseURL, userID, shadowAccount.ID, depositCount, amount); err != nil {
				log.Fatalln(err)
				return
			}
		}()
	}

	wg.Wait()

	// Check shadow account balance
	shadowAccount, err = cli.GetAccount(shadowAccount.ID)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	if wantBalance := int(beforeBalance) - 1*depositCount*userCount*int(amount); int(shadowAccount.Balance) != wantBalance {
		t.Errorf("want %d , got %d", wantBalance, shadowAccount.Balance)
		t.FailNow()
	}

	if err = CheckHistory(cli, "admin", shadowAccount.ID, shadowAccount.Balance); err != nil {
		t.Error(err)
		t.FailNow()
	}
}
