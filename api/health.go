package api

import (
	"goat/services/health"

	// "goat/log"
	"net/http"
)

type HealthEndpoint struct {
	service *health.Service
	// logger  log.Logger
}

func NewHealthEndpoint(service *health.Service) *HealthEndpoint {
	return &HealthEndpoint{
		service: service,
		// logger:  logger.WithPrefix("api.health"),
	}
}

func (e *HealthEndpoint) GetHealth(w http.ResponseWriter, _ *http.Request) {
	var status int
	healthResponse := e.service.GetHealth()

	switch healthResponse.State {
	case health.StateStarting:
		status = http.StatusAccepted
	case health.StateRunning:
		status = http.StatusOK
	case health.StateStopping:
		status = http.StatusExpectationFailed
	default:
		status = http.StatusInternalServerError
	}
	// 	respond(w, e.logger, status, "", healthResponse)
	respond(w, status, "", healthResponse)
}
