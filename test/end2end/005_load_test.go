package end2end

//
//import (
//	"fmt"
//	"github.com/sinameshkini/fingo/internal/models"
//	"github.com/sinameshkini/fingo/pkg/enums"
//	"github.com/sinameshkini/fingo/test"
//	"testing"
//	"time"
//)
//
//// ID:			TS005_Load
//// Scenario:	High load request (TPS calculate)
//
//func Test_TS005_Load(t *testing.T) {
//	var (
//		userCount                  = 2
//		depositCount               = 20
//		amount       models.Amount = 10
//	)
//
//	cli, baseURL, err := test.Setup()
//	if err != nil {
//		t.Error(err.Error())
//		t.FailNow()
//	}
//
//	// Get shadow account for deposit transaction
//	shadowAccount, err := GetAccount(cli, "admin", enums.ACCOUNTTYPESHADOW)
//	if err != nil {
//		t.Error(err.Error())
//		t.FailNow()
//	}
//
//	// Concurrent account creation
//	for i := 0; i < userCount; i++ {
//		userID := MakeUserID(i)
//		go func() {
//			if err = NormalActor(baseURL, userID, shadowAccount.ID, depositCount, amount); err != nil {
//				t.Error(err.Error())
//				t.FailNow()
//			}
//		}()
//	}
//
//	time.Sleep(time.Second * 3)
//
//	// Check shadow account balance
//	shadowAccount, err = cli.GetAccount(shadowAccount.ID)
//	if err != nil {
//		t.Error(err.Error())
//		t.FailNow()
//	}
//
//	if wantBalance := -1 * depositCount * userCount * int(amount); int(shadowAccount.Balance) != wantBalance {
//		t.Error("want", wantBalance, "got", shadowAccount.Balance)
//	}
//}
//
//func MakeUserID(id int) string {
//	return fmt.Sprintf("user%d", id)
//}
