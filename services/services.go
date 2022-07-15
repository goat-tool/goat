package services

import (
	"goat/services/health"
	"goat/services/root"
	"goat/services/test"
	"goat/services/user"
)

// Services holds all available services
type Services struct {
	Root   *root.Service
	Health *health.Service
	User   *user.Service
	Test   *test.Service
}
