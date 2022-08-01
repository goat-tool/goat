package services

import (
	"goat/services/example"
	"goat/services/health"
	"goat/services/root"
	"goat/services/user"
)

// Services holds all available services
type Services struct {
	Root    *root.Service
	Health  *health.Service
	User    *user.Service
	Example *example.Service
}
