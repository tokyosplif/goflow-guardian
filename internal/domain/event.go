package domain

import "time"

const ReasonRateLimitExceeded = "rate_limit_exceeded"

type Violation struct {
	Key       string    `json:"key"`
	Reason    string    `json:"reason"`
	Timestamp time.Time `json:"timestamp"`
}
