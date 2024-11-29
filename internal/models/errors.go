package models

import "errors"

var (
	ErrAccountTypeInvalid = errors.New("account type is invalid")
	ErrCurrencyInvalid    = errors.New("currency is invalid")
	ErrInternal           = errors.New("internal error, try again later")
	ErrNotFound           = errors.New("not found")
	ErrPermissionDenied   = errors.New("permission denied")
	ErrNotEnoughBalance   = errors.New("not enough balance")
	ErrInvalidRequest     = errors.New("invalid request")
)
