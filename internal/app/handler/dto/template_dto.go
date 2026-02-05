package dto

type HealthStatus string

const (
	HealthStatusUp   HealthStatus = "UP"
	HealthStatusDown HealthStatus = "DOWN"
)

type HealthResponse struct {
	Status HealthStatus `json:"status"`
}