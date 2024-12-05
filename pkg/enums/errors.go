package enums

import (
	"errors"
	"net/http"
)

var (
	ErrAccountTypeInvalid = errors.New("account type is invalid")
	ErrCurrencyInvalid    = errors.New("currency is invalid")
	ErrInternal           = errors.New("internal error, try again later")
	ErrNotFound           = errors.New("not found")
	ErrPermissionDenied   = errors.New("permission denied")
	ErrNotEnoughBalance   = errors.New("not enough balance")
	ErrInvalidRequest     = errors.New("invalid request")
	ErrToManyRequests     = errors.New("too many requests")
)

var ErrCode = map[error]int{
	ErrAccountTypeInvalid: 1001,
	ErrCurrencyInvalid:    1002,
	ErrInternal:           1003,
	ErrNotFound:           1004,
	ErrPermissionDenied:   1005,
	ErrNotEnoughBalance:   1006,
	ErrInvalidRequest:     1007,
	ErrToManyRequests:     1008,
}

var ErrHTTPCode = map[error]int{
	ErrInvalidRequest:   http.StatusBadRequest,
	ErrNotFound:         http.StatusNotFound,
	ErrPermissionDenied: http.StatusForbidden,
	ErrToManyRequests:   http.StatusTooManyRequests,
}
