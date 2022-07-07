package core

import (
	"goat/services"
	"goat/services/health"
	"goat/services/user"
)

func (c *Core) setupServices() {

	c.Log.Debug().Msg("Todo: setupServices()")

	healthService := health.NewService(&c.state)
	userService := user.NewService()

	c.services = &services.Services{
		Health: healthService,
		User:   userService,
	}

	// healthService := health.NewService(c.Log, &c.state)
	// userService := user.NewService(c.Log, user.NewStore(c.Log, c.Config, c.Database))

	// c.Services = &services.Services{
	// 	Health: healthService,
	// 	User:   userService,
	// }
}
