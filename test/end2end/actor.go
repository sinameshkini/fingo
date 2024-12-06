package end2end

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/sinameshkini/fingo/pkg/endpoint"
	"github.com/sinameshkini/fingo/pkg/enums"
	"github.com/sinameshkini/fingo/pkg/sdk"
	"github.com/sinameshkini/microkit/models"
	"github.com/sinameshkini/microkit/pkg/utils"
)

func CreateAccount(cli *sdk.Client, userID, accountType, name string) (resp *endpoint.AccountResponse, err error) {
	currencies, err := cli.GetCurrencies()
	if err != nil {
		return
	}

	accountTypes, err := cli.GetAccountTypes()
	if err != nil {
		return
	}

	if len(currencies) == 0 || len(accountTypes) == 0 {
		return nil, errors.New("no currencies or accountTypes")
	}

	account, err := cli.CreateAccount(endpoint.CreateAccount{
		UserID:        userID,
		AccountTypeID: accountType,
		CurrencyID:    currencies[0].ID,
		Name:          name,
	})
	if err != nil {
		return
	}

	return cli.GetAccount(account.ID)
}

func CreateAccountIfNotExist(cli *sdk.Client, userID, accountType, name string) (resp *endpoint.AccountResponse, err error) {
	if resp, err = GetAccount(cli, userID, accountType); err == nil {
		return
	}

	currencies, err := cli.GetCurrencies()
	if err != nil {
		return
	}

	accountTypes, err := cli.GetAccountTypes()
	if err != nil {
		return
	}

	if len(currencies) == 0 || len(accountTypes) == 0 {
		return nil, errors.New("no currencies or accountTypes")
	}

	account, err := cli.CreateAccount(endpoint.CreateAccount{
		UserID:        userID,
		AccountTypeID: accountType,
		CurrencyID:    currencies[0].ID,
		Name:          name,
	})
	if err != nil {
		return
	}

	return cli.GetAccount(account.ID)
}

func GetAccount(cli *sdk.Client, userID, accountType string) (resp *endpoint.AccountResponse, err error) {
	adminAccounts, err := cli.GetAccounts(userID)
	if err != nil {
		return
	}

	for _, aa := range adminAccounts {
		if aa.AccountType.ID == accountType {
			return aa, nil
		}
	}

	return nil, enums.ErrNotFound
}

func CheckHistory(cli *sdk.Client, userID, accountID string, fromBalance models.Amount) (err error) {
	var (
		page        int64 = 1
		balanceGage       = fromBalance
	)

	for {
		historyResp, err := cli.History(endpoint.HistoryRequest{
			PaginationRequest: models.PaginationRequest{Page: page},
			UserID:            userID,
			AccountID:         accountID,
		})
		if err != nil {
			return err
		}

		for _, t := range historyResp.Transactions {
			if balanceGage != t.Balance {
				return fmt.Errorf("invalid transaction balance, got: %d, want: %d", t.Balance, balanceGage)
			}

			balanceGage += t.Amount * -1
		}

		if !historyResp.Meta.HasNext {
			if len(historyResp.Transactions) != 0 {
				if first := historyResp.Transactions[len(historyResp.Transactions)-1]; first.Balance-first.Amount != 0 {
					return fmt.Errorf("invalid transaction balance, got: %d, want: 0", first.Balance-first.Amount)
				}
			}
			break
		}

		page++
	}

	return
}

func NormalActor(baseURL, userID, shadow string, cnt int, amount models.Amount) (err error) {
	cli := sdk.New(baseURL, true)

	account, err := CreateAccountIfNotExist(cli, userID, enums.ACCOUNTTYPEWALLET, fmt.Sprintf("%s-%s", userID, enums.ACCOUNTTYPEWALLET))
	if err != nil {
		return
	}

	beforeBalance := account.Balance

	for i := 0; i < cnt; i++ {
		depositTxn, err := cli.Transfer(endpoint.TransferRequest{
			UserID:          "admin",
			Type:            enums.Deposit,
			OrderID:         uuid.NewString(),
			DebitAccountID:  shadow,
			CreditAccountID: account.ID,
			RawAmount:       amount,
			TotalAmount:     amount,
			Description:     fmt.Sprintf("%s deposit %d", userID, i),
		})
		if err != nil {
			return err
		}

		utils.PrintJson(depositTxn)
	}

	account, err = GetAccount(cli, userID, enums.ACCOUNTTYPEWALLET)
	if err != nil {
		return
	}

	if account.Balance != beforeBalance+amount*models.Amount(cnt) {
		return errors.New("not enough balance")
	}

	if err = CheckHistory(cli, userID, account.ID, account.Balance); err != nil {
		return
	}

	return nil
}

func MakeUserID(id int) string {
	return fmt.Sprintf("user%d", id)
}
