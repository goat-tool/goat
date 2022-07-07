package core

// "net/http"

func (c *Core) setupRouter() {

	c.Log.Debug().Msg("Todo: setupRouter()")

	// logger := c.Log.WithPrefix("core.router")
	// middlewares := setupMiddleware(c.Router, logger)

	// // register middlewares
	// middlewares = loggingMiddleware(middlewares, logger)

	// c.Handler = middlewares
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
