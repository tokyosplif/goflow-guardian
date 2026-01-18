package domain

import "errors"

var (
	ErrLimitExceeded = errors.New("rate limit exceeded")
)
