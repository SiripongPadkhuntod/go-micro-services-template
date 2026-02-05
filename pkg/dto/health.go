package dto

const (
	HealthStatusUp   = "up"
	HealthStatusDown = "down"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
	Version string `json:"version"`
}
