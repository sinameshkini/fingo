package database

import (
	"github.com/labstack/gommon/log"
	"github.com/sinameshkini/fingo/internal/models"
	"gorm.io/gorm"
)

func Drop(db *gorm.DB) (err error) {
	for _, t := range tables {
		if err = db.Migrator().DropTable(t); err != nil {
			log.Error(err.Error())
		}
	}

	return nil
}

func Migrate(db *gorm.DB) (err error) {
	if err = db.AutoMigrate(tables...); err != nil {
		return err
	}

	return nil
}

var (
	minAmount   models.Amount = 10000
	maxAmount   models.Amount = 10000000
	defaultCode               = models.CodeP{
		FeeType:                 models.FeeTypePointer(models.FeeActual),
		FeeValue:                models.AmountPointer(0),
		MinAmountPerTransaction: models.AmountPointer(minAmount),
		MaxAmountPerTransaction: models.AmountPointer(maxAmount),
		MaxAmountPerDay:         models.AmountPointer(maxAmount * 3),
		MaxCountPerDay:          models.IntPointer(3),
	}

	defaultCodeWithFee = models.CodeP{
		FeeType:                 models.FeeTypePointer(models.FeeActual),
		FeeValue:                models.AmountPointer(1000),
		MinAmountPerTransaction: models.AmountPointer(minAmount),
		MaxAmountPerTransaction: models.AmountPointer(maxAmount),
		MaxAmountPerDay:         models.AmountPointer(maxAmount * 3),
		MaxCountPerDay:          models.IntPointer(3),
	}
)

func Seed(db *gorm.DB) (err error) {

	//	accountTypes
	accountTypes := []*models.AccountType{
		//{
		//	ID:          models.ACCOUNTTYPEGL,
		//	Name:        "GL",
		//	Description: "general ledger",
		//},
		{
			ID:          models.ACCOUNTTYPEWALLET,
			Name:        "wallet",
			Description: "Wallet",
		},
		{
			ID:          models.ACCOUNTTYPESHADOW,
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
		//{
		//	ID:          models.ACCOUNTTYPEFEE,
		//	Name:        "fee account",
		//	Description: "Fee Account",
		//},
	}

	for _, a := range accountTypes {
		if err = db.FirstOrCreate(&a).Error; err != nil {
			return
		}
	}

	// user_group 0-9
	// account_type 20-39
	//

	policies := []*models.Policy{
		{
			EntityType: "user_group",
			EntityID:   "",
			Settings: models.SettingsP{
				Limits: &models.LimitsP{
					MinBalance:       models.AmountPointer(0),
					MaxBalance:       models.AmountPointer(maxAmount),
					NumberOfAccounts: map[string]uint{models.ACCOUNTTYPEWALLET: 1},
				},
				DefaultAccountTypeID: models.StringPointer(models.ACCOUNTTYPEWALLET),
			},
			Priority: 0,
			IsEnable: true,
		},
		{
			EntityType: "account_type",
			EntityID:   models.ACCOUNTTYPEWALLET,
			Settings: models.SettingsP{
				Limits: &models.LimitsP{
					MinBalance: models.AmountPointer(0),
					MaxBalance: models.AmountPointer(maxAmount),
				},
				Codes: map[models.ProcessCode]models.CodeP{
					models.CodeDepositCredit:  defaultCode,
					models.CodePurchaseDebit:  defaultCode,
					models.CodeTransferDebit:  defaultCodeWithFee,
					models.CodeTransferCredit: defaultCode,
					models.CodeWithdrawDebit:  defaultCodeWithFee,
				},
			},
			Priority: 20,
			IsEnable: true,
		},
		{
			EntityType: "account_type",
			EntityID:   models.ACCOUNTTYPESHADOW,
			Settings: models.SettingsP{
				Limits: &models.LimitsP{
					MinBalance:       models.AmountPointer(-10000000000),
					MaxBalance:       models.AmountPointer(0),
					NumberOfAccounts: map[string]uint{models.ACCOUNTTYPESHADOW: 1},
				},
				Codes: map[models.ProcessCode]models.CodeP{
					models.CodeDepositDebit: defaultCode,
				},
			},
			Priority: 35,
			IsEnable: true,
		},
		{
			EntityType: "account_type",
			EntityID:   models.ACCOUNTTYPETERMINAL,
			Settings: models.SettingsP{
				Limits: &models.LimitsP{
					MinBalance:       models.AmountPointer(0),
					MaxBalance:       models.AmountPointer(maxAmount),
					NumberOfAccounts: map[string]uint{models.ACCOUNTTYPETERMINAL: 1},
				},
				Codes: map[models.ProcessCode]models.CodeP{
					models.CodePurchaseCredit: defaultCode,
				},
			},
			Priority: 30,
			IsEnable: true,
		},
		{
			EntityType: "account_type",
			EntityID:   models.ACCOUNTTYPEEXTERNALACCOUNT,
			Settings: models.SettingsP{
				Limits: &models.LimitsP{
					MinBalance:       models.AmountPointer(0),
					MaxBalance:       models.AmountPointer(maxAmount),
					NumberOfAccounts: map[string]uint{models.ACCOUNTTYPEEXTERNALACCOUNT: 1},
				},
				Codes: map[models.ProcessCode]models.CodeP{
					models.CodeWithdrawCredit: defaultCode,
				},
			},
			Priority: 25,
			IsEnable: true,
		},
		{
			EntityType: "account_type",
			EntityID:   models.ACCOUNTTYPEFEE,
			Settings: models.SettingsP{
				Limits: &models.LimitsP{
					MinBalance:       models.AmountPointer(0),
					MaxBalance:       models.AmountPointer(maxAmount),
					NumberOfAccounts: map[string]uint{models.ACCOUNTTYPEFEE: 1},
				},
				Codes: map[models.ProcessCode]models.CodeP{
					models.CodeWithdrawCredit: defaultCode,
				},
			},
			Priority: 32,
			IsEnable: true,
		},
		{
			EntityType: "account_type",
			EntityID:   "9",
			Settings: models.SettingsP{
				Limits: &models.LimitsP{
					MinBalance:       models.AmountPointer(0),
					MaxBalance:       models.AmountPointer(maxAmount),
					NumberOfAccounts: map[string]uint{models.ACCOUNTTYPEWALLET: 1},
				},
				Codes: map[models.ProcessCode]models.CodeP{
					models.CodeDepositCredit: defaultCode,
					models.CodePurchaseDebit: defaultCode,
				},
			},
			Priority: 30,
			IsEnable: true,
		},
	}

	for _, p := range policies {
		if err = db.Create(&p).Error; err != nil {
			return
		}
	}

	//	currencies
	currencies := []*models.Currency{
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
	accounts := []*models.Account{
		//{
		//	Model:         models.Model{ID: models.ID2},
		//	AccountTypeID: models.ACCOUNTTYPETERMINAL,
		//	CurrencyID:    1,
		//	Name:          "service provider terminal",
		//	IsEnable:      true,
		//},
		//{
		//	Model:         models.Model{ID: models.ID6},
		//	AccountTypeID: models.ACCOUNTTYPETERMINAL,
		//	CurrencyID:    1,
		//	Name:          "wallet fee terminal",
		//	IsEnable:      false,
		//},
		{
			Model:         models.Model{ID: models.ID3},
			AccountTypeID: models.ACCOUNTTYPESHADOW,
			CurrencyID:    1,
			Name:          "payment gateway shadow",
			IsEnable:      true,
			UserID:        "admin",
		},
		//{
		//	Model:         models.Model{ID: models.ID4},
		//	AccountTypeID: models.ACCOUNTTYPEFEE,
		//	CurrencyID:    1,
		//	Name:          "fee",
		//	IsEnable:      true,
		//},
	}

	for _, a := range accounts {
		if err = db.FirstOrCreate(&a).Error; err != nil {
			return
		}
	}

	return nil
}
