package core

func (c *Core) setupApi() {

	c.Log.Debug().Msg("Todo: setupApi()")

	// c.Api = &api.Api{
	// 	Health: api.NewHealthEndpoint(c.Services.Health, c.Log),
	// 	User:   api.NewUserEndpoint(c.Log, c.Translator, c.Validate, c.Services.User),
	// }

	// c.Api.Setup(c.Router)
}
