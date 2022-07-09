package core

import (
	"net/http"
	"time"
)

type logEntry struct {
	ReceivedTime       time.Time
	RequestMethod      string
	RequestURL         string
	RequestHeaderSize  int64
	RequestBodySize    int64
	UserAgent          string
	Referer            string
	Proto              string
	RemoteIP           string
	ServerIP           string
	Status             int
	ResponseHeaderSize int64
	ResponseBodySize   int64
	Latency            time.Duration
}

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

		h.ServeHTTP(w, r)

		c.Log.Debug().
			Str("agent", r.UserAgent()).
			Str("referer", r.Referer()).
			Str("proto", r.Proto).
			Str("remote_addr", r.RemoteAddr).
			Str("method", r.Method).
			Str("url", r.URL.String()).
			Msg("Router")

	})
}
