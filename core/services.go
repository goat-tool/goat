package core

import (
	"goat/services"
	"goat/services/example"
	"goat/services/health"
	"goat/services/root"
	"goat/services/user"
)

func (c *Core) newServices() *services.Services {

	c.Log.Info().Msg("Setup services")
	rootService := root.NewService()
	healthService := health.NewService(&c.state)
	userService := user.NewService(c.Log, c.Conf, c.Database)
	exampleService := example.NewService(c.Log, c.Conf, c.Database)

	return &services.Services{
		Root:    rootService,
		Health:  healthService,
		User:    userService,
		Example: exampleService,
	}
}
