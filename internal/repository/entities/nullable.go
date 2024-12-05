package entities

import "github.com/sinameshkini/fingo/pkg/types"

func AmountPointer(amount types.Amount) *types.Amount {
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
