package service

import (
	"AISwimCoachAppBackend/src/vo"
	"math/rand"
)

// HealthCheckHandler is a simple handler to verify the API is working
func HealthCheck() vo.HealthCheck {
	status := "OK"
	message := "API is healthy"

	if rand.Intn(2) == 0 {
		status = "BAD"
		message = "API is unhealthy"
	}

	return vo.HealthCheck{Status: status, Message: message}
}
