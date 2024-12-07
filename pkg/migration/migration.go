package migration

import (
	"github.com/sinameshkini/fingo/internal/repository/entities"
	"github.com/sinameshkini/fingo/pkg/enums"
	"github.com/sinameshkini/microkit/models"
	"gorm.io/gorm"
)

var (
	minAmount   models.Amount = 10000
	maxAmount   models.Amount = 10000000
	defaultCode               = entities.CodeP{
		FeeType:                 entities.FeeTypePointer(entities.FeeActual),
		FeeValue:                entities.AmountPointer(0),
		MinAmountPerTransaction: entities.AmountPointer(minAmount),
		MaxAmountPerTransaction: entities.AmountPointer(maxAmount),
		MaxAmountPerDay:         entities.AmountPointer(maxAmount * 3),
		MaxCountPerDay:          entities.IntPointer(3),
	}

	defaultCodeWithFee = entities.CodeP{
		FeeType:                 entities.FeeTypePointer(entities.FeeActual),
		FeeValue:                entities.AmountPointer(1000),
		MinAmountPerTransaction: entities.AmountPointer(minAmount),
		MaxAmountPerTransaction: entities.AmountPointer(maxAmount),
		MaxAmountPerDay:         entities.AmountPointer(maxAmount * 3),
		MaxCountPerDay:          entities.IntPointer(3),
	}
)

func Seed(db *gorm.DB) (err error) {

	//	accountTypes
	accountTypes := []*entities.AccountType{
		{
			ID:          enums.ACCOUNTTYPEWALLET,
			Name:        "wallet",
			Description: "Wallet",
		},
		{
			ID:          enums.ACCOUNTTYPESHADOW,
			Name:        "shadow",
			Description: "Payment Gateway",
		},
		//{
		//	ID:          models.ACCOUNTTYPETERMINAL,
		//	Name:        "terminal",
		//	Description: "Merchant",
		//},
		//{
		//	ID:          models.ACCOUNTTYPELOAN,
		//	Name:        "loan",
		//	Description: "Loan or Credit",
		//},
		//{
		//	ID:          models.ACCOUNTTYPEINSTALLMENT,
		//	Name:        "installment",
		//	Description: "Installment",
		//},
		//{
		//	ID:          models.ACCOUNTTYPEEXTERNALACCOUNT,
		//	Name:        "bank account",
		//	Description: "Bank Account",
		//},
		{
			ID:          enums.ACCOUNTTYPEFEE,
			Name:        "fee account",
			Description: "Fee Account",
		},
	}

	for _, a := range accountTypes {
		if err = db.FirstOrCreate(&a).Error; err != nil {
			return
		}
	}

	policies := []*entities.Policy{
		{
			EntityType: "user_group",
			EntityID:   "",
			Settings: entities.SettingsP{
				Limits: map[string]entities.LimitsP{
					enums.ACCOUNTTYPEWALLET: {
						MinBalance:       entities.AmountPointer(0),
						MaxBalance:       entities.AmountPointer(maxAmount),
						NumberOfAccounts: entities.IntPointer(1),
					},
				},
				DefaultAccountTypeID: entities.StringPointer(enums.ACCOUNTTYPEWALLET),
			},
			Priority: 0,
			IsEnable: true,
		},
		{
			EntityType: "account_type",
			EntityID:   enums.ACCOUNTTYPEWALLET,
			Settings: entities.SettingsP{
				Limits: map[string]entities.LimitsP{
					enums.ACCOUNTTYPEWALLET: {
						MinBalance: entities.AmountPointer(0),
						MaxBalance: entities.AmountPointer(maxAmount),
					},
				},
				Codes: map[enums.ProcessCode]entities.CodeP{
					enums.CodeDepositCredit:  defaultCode,
					enums.CodePurchaseDebit:  defaultCode,
					enums.CodeTransferDebit:  defaultCodeWithFee,
					enums.CodeTransferCredit: defaultCode,
					enums.CodeWithdrawDebit:  defaultCodeWithFee,
				},
			},
			Priority: 20,
			IsEnable: true,
		},
		{
			EntityType: "account_type",
			EntityID:   enums.ACCOUNTTYPESHADOW,
			Settings: entities.SettingsP{
				Limits: map[string]entities.LimitsP{
					enums.ACCOUNTTYPESHADOW: {
						MinBalance:       entities.AmountPointer(-10000000000),
						MaxBalance:       entities.AmountPointer(0),
						NumberOfAccounts: entities.IntPointer(1),
					},
				},
				Codes: map[enums.ProcessCode]entities.CodeP{
					enums.CodeDepositDebit: defaultCode,
				},
			},
			Priority: 35,
			IsEnable: true,
		},
		{
			EntityType: "account_type",
			EntityID:   enums.ACCOUNTTYPETERMINAL,
			Settings: entities.SettingsP{
				Limits: map[string]entities.LimitsP{
					enums.ACCOUNTTYPETERMINAL: {
						MinBalance:       entities.AmountPointer(0),
						MaxBalance:       entities.AmountPointer(maxAmount),
						NumberOfAccounts: entities.IntPointer(1),
					},
				},
				Codes: map[enums.ProcessCode]entities.CodeP{
					enums.CodePurchaseCredit: defaultCode,
				},
			},
			Priority: 30,
			IsEnable: true,
		},
		{
			EntityType: "account_type",
			EntityID:   enums.ACCOUNTTYPEEXTERNALACCOUNT,
			Settings: entities.SettingsP{
				Limits: map[string]entities.LimitsP{
					enums.ACCOUNTTYPEEXTERNALACCOUNT: {
						MinBalance:       entities.AmountPointer(0),
						MaxBalance:       entities.AmountPointer(maxAmount),
						NumberOfAccounts: entities.IntPointer(1),
					},
				},
				Codes: map[enums.ProcessCode]entities.CodeP{
					enums.CodeWithdrawCredit: defaultCode,
				},
			},
			Priority: 25,
			IsEnable: true,
		},
		{
			EntityType: "account_type",
			EntityID:   enums.ACCOUNTTYPEFEE,
			Settings: entities.SettingsP{
				Limits: map[string]entities.LimitsP{
					enums.ACCOUNTTYPEFEE: {
						MinBalance:       entities.AmountPointer(0),
						MaxBalance:       entities.AmountPointer(maxAmount),
						NumberOfAccounts: entities.IntPointer(1),
					},
				},
				Codes: map[enums.ProcessCode]entities.CodeP{
					enums.CodeWithdrawCredit: defaultCode,
				},
			},
			Priority: 32,
			IsEnable: true,
		},
	}

	for _, p := range policies {
		if err = db.Create(&p).Error; err != nil {
			return
		}
	}

	//	currencies
	currencies := []*entities.Currency{
		{
			ID:        1,
			Symbol:    "USD",
			IsEnable:  true,
			IsDefault: true,
		},
	}

	for _, c := range currencies {
		if err = db.FirstOrCreate(&c).Error; err != nil {
			return
		}
	}

	//	accounts
	accounts := []*entities.Account{
		{
			AccountTypeID: enums.ACCOUNTTYPESHADOW,
			CurrencyID:    1,
			Name:          "payment gateway shadow",
			IsEnable:      true,
			UserID:        "admin",
		},
		{
			AccountTypeID: enums.ACCOUNTTYPEFEE,
			CurrencyID:    1,
			Name:          "fee",
			IsEnable:      true,
			UserID:        "admin",
		},
	}

	for _, a := range accounts {
		if err = db.Create(&a).Error; err != nil {
			return
		}
	}

	return nil
}
