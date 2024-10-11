package errs

import "errors"

var (
	ErrCardNumberInvalid      = errors.New("invalid card number")
	ErrExpirationMonthInvalid = errors.New("invalid expiration month")
	ErrExpiredMonth           = errors.New("expired month")
	ErrExpiredYear            = errors.New("expired year")
)
