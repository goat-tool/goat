package core

import (
	"net/http"
)

func (c *Core) setupRouter() {

	c.Log.Info().Msg("Setup router")

	middlewares := setupMiddleware(c.router, c)

	// // register middlewares
	middlewares = loggingMiddleware(middlewares, c)

	c.handler = middlewares
}

func setupMiddleware(h http.Handler, c *Core) http.Handler {
	c.Log.Info().Msg("Setup router middlewares")
	return h
}

func loggingMiddleware(h http.Handler, c *Core) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.Log.Debug().Str("method", r.Method).Str("addr", r.RemoteAddr).Str("request", r.RequestURI).Msg("Router")
		h.ServeHTTP(w, r)
	})
}
