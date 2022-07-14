package core

import (
	"goat/services"
	"goat/services/health"
	"goat/services/test"
	"goat/services/user"
)

func (c *Core) newServices() {

	c.Log.Info().Msg("Setup services")

	healthService := health.NewService(&c.state)
	userService := user.NewService()
	testService := test.NewService(c.Log, c.Conf, c.Database)

	c.services = &services.Services{
		Health: healthService,
		User:   userService,
		Test:   testService,
	}
}
