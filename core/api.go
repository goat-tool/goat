package core

import "goat/api"

func (c *Core) newApi() {

	c.Log.Info().Msg("Setup api")

	c.api = &api.Api{
		Health: api.NewHealthEndpoint(c.Log, c.services.Health),
		User:   api.NewUserEndpoint(c.Log, c.services.User),
		Test:   api.NewTestEndpoint(c.Log, c.services.Test),
	}

	c.api.Setup(c.router)

}
