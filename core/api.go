package core

import "goat/api"

func (c *Core) setupApi() {

	c.Log.Info().Msg("Setup api")

	c.api = &api.Api{
		Health: api.NewHealthEndpoint(c.services.Health),
		User:   api.NewUserEndpoint(c.services.User),
	}

	c.api.Setup(c.router)

	// c.Api = &api.Api{
	// 	Health: api.NewHealthEndpoint(c.Services.Health, c.Log),
	// 	User:   api.NewUserEndpoint(c.Log, c.Translator, c.Validate, c.Services.User),
	// }

	// c.Api.Setup(c.Router)
}
