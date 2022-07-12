package services

import (
	"goat/services/health"
	"goat/services/test"
	"goat/services/user"
)

// Services holds all available services
type Services struct {
	Health *health.Service
	User   *user.Service
	Test   *test.Service
}
