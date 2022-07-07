package services

import (
	"goat/services/health"
	"goat/services/user"
)

// Services holds all available services
type Services struct {
	Health *health.Service
	User   *user.Service
}
