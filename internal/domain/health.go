package domain

const (
	StatusOK   = "OK"
	StatusDown = "DOWN"
)

type HealthStatus struct {
	Status     string            `json:"status"`
	Components map[string]string `json:"components"`
}
