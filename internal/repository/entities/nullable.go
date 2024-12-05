package entities

import "github.com/sinameshkini/microkit/models"

func AmountPointer(amount models.Amount) *models.Amount {
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
