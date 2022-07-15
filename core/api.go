package core

import "goat/api"

func (c *Core) newApi() *api.Api {

	c.Log.Info().Msg("Setup api")

	api := &api.Api{
		Root:   api.NewRootEndpoint(c.Log, c.services.Root),
		Health: api.NewHealthEndpoint(c.Log, c.services.Health),
		User:   api.NewUserEndpoint(c.Log, c.services.User),
		Test:   api.NewTestEndpoint(c.Log, c.services.Test),
	}

	api.New(c.router)

	return api
}
