package core

import (
	"goat/services"
	"goat/services/health"
	"goat/services/test"
	"goat/services/user"
)

func (c *Core) newServices() *services.Services {

	c.Log.Info().Msg("Setup services")

	healthService := health.NewService(&c.state)
	userService := user.NewService(c.Log, c.Conf, c.Database)
	testService := test.NewService(c.Log, c.Conf, c.Database)

	return &services.Services{
		Health: healthService,
		User:   userService,
		Test:   testService,
	}
}
