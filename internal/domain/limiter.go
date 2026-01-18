package domain

import "time"

const (
	StatusAllowed  = "allowed"
	StatusRejected = "rejected"
)

type Limit struct {
	Requests int
	Window   time.Duration
}
