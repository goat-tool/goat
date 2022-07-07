package core

import (
	"fmt"
	"net/http"
)

func (c *Core) setupRouter() {

	c.Log.Debug().Msg("Todo: setupRouter()")

	// logger := c.Log.WithPrefix("core.router")
	middlewares := setupMiddleware(c.router)

	// // register middlewares
	middlewares = loggingMiddleware(middlewares)

	c.handler = middlewares
}

func setupMiddleware(h http.Handler) http.Handler {
	// logger.Debugln("setting up router middlewares")
	return h
}

func loggingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// logger.WithFields(log.Fields{"method": r.Method, "addr": r.RemoteAddr}).Debugln("request:", r.RequestURI)
		fmt.Println(r.RequestURI)
		h.ServeHTTP(w, r)
	})
}

// func setupMiddleware(h http.Handler, logger log.Logger) http.Handler {
// 	logger.Debugln("setting up router middlewares")
// 	return h
// }

// func loggingMiddleware(h http.Handler, logger log.Logger) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		logger.WithFields(log.Fields{"method": r.Method, "addr": r.RemoteAddr}).Debugln("request:", r.RequestURI)
// 		h.ServeHTTP(w, r)
// 	})
// }
