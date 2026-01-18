package domain

import "errors"

var (
	ErrInternal      = errors.New("internal error")
	ErrLimitExceeded = errors.New("rate limit exceeded")
)
