package models

func AmountPointer(amount Amount) *Amount {
	return &amount
}

func StringPointer(str string) *string {
	return &str
}

func IntPointer(i int) *int {
	return &i
}

func FeeTypePointer(t FeeType) *FeeType {
	return &t
}
