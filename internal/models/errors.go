package models

import "errors"

var (
	ErrAccountTypeInvalid = errors.New("account type is invalid")
	ErrCurrencyInvalid    = errors.New("currency is invalid")
	ErrInternal           = errors.New("internal error, try again later")
)
