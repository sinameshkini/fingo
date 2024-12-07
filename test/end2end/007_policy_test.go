package end2end

import (
	"github.com/sinameshkini/fingo/pkg/endpoint"
	"github.com/sinameshkini/fingo/test"
	"github.com/sinameshkini/microkit/pkg/utils"
	"testing"
)

// ID:			TS007_History
// Scenario:	Policy

func Test_Policy(t *testing.T) {
	cli, _, err := test.Setup()
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	settingsResp, err := cli.GetSettings(endpoint.GetSettingsRequest{})
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	utils.PrintJson(settingsResp)

}
