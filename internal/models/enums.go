package models

type TransactionType string

const (
	Deposit    TransactionType = "deposit"
	Purchase   TransactionType = "purchase"
	Transfer   TransactionType = "transfer"
	Withdrawal TransactionType = "withdrawal"
	Reverse    TransactionType = "reverse"
	PayLoan    TransactionType = "pay_loan"
	PayoffLoan TransactionType = "payoff_loan"
	Unknown    TransactionType = "unknown"
)

func (t TransactionType) Label() string {
	// TODO move to response map
	switch t {
	case Deposit:
		return "افزایش موجودی"
	case Purchase:
		return "خرید"
	case Transfer:
		return "انتقال به حساب دیگر"
	case Withdrawal:
		return "برداشت از حساب"
	case Reverse:
		return "سند یازگشت تراکنش"
	case PayLoan:
		return "اعطای وام / اعتبار"
	case PayoffLoan:
		return "یازپرداخت قسط / بدهی"
	}

	return ""
}

type DocumentType string

const (
	Credit DocumentType = "credit"
	Debit  DocumentType = "debit"
)

func (d DocumentType) Reverse() DocumentType {
	if d == Credit {
		return Debit
	}

	return Credit
}

type ProcessCode string

const (
	CodeDepositDebit       ProcessCode = "deposit_debit"
	CodeDepositCredit      ProcessCode = "deposit_credit"
	CodePurchaseDebit      ProcessCode = "purchase_debit"
	CodePurchaseCredit     ProcessCode = "purchase_credit"
	CodeTransferDebit      ProcessCode = "transfer_debit"
	CodeTransferCredit     ProcessCode = "transfer_credit"
	CodeWithdrawDebit      ProcessCode = "withdrawal_debit"
	CodeWithdrawCredit     ProcessCode = "withdrawal_credit"
	CodeReverseDebit       ProcessCode = "reverse_debit"
	CodeReverseCredit      ProcessCode = "reverse_credit"
	CodeLoanOrCreditDebit  ProcessCode = "loan_debit"
	CodeLoanOrCreditCredit ProcessCode = "loan_credit"
)

func (pc ProcessCode) TransactionType() TransactionType {
	switch pc {
	case CodeDepositDebit:
	case CodeDepositCredit:
		return Deposit

	case CodePurchaseDebit:
	case CodePurchaseCredit:
		return Purchase

	case CodeTransferDebit:
	case CodeTransferCredit:
		return Transfer

	case CodeWithdrawDebit:
	case CodeWithdrawCredit:
		return Withdrawal

	case CodeReverseDebit:
	case CodeReverseCredit:
		return Reverse

	case CodeLoanOrCreditDebit:
	case CodeLoanOrCreditCredit:
		return PayLoan

	}

	return Unknown
}

var (
	ID1 ID = 1273444730903568384
	ID2 ID = 1273444735114649600
	ID3 ID = 1273444739325730816
	ID4 ID = 1273444743536812032
	ID5 ID = 1273444747747893248
	ID6 ID = 1273444760381136896
	ID7 ID = 1273444768803299328
)

const (
	ACCOUNTTYPEGL              = "1"
	ACCOUNTTYPEWALLET          = "2"
	ACCOUNTTYPESHADOW          = "3"
	ACCOUNTTYPETERMINAL        = "4"
	ACCOUNTTYPELOAN            = "5"
	ACCOUNTTYPEINSTALLMENT     = "6"
	ACCOUNTTYPEEXTERNALACCOUNT = "7"
	ACCOUNTTYPEFEE             = "8"
)
