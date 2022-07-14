package core

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (c *Core) newRouter() {

	c.Log.Info().Msg("Setup router")

	c.router = mux.NewRouter()
	middlewares := setupMiddleware(c.router, c)

	// register middlewares
	middlewares = jsonMiddleware(middlewares)
	middlewares = loggingMiddleware(middlewares, c)

	c.handler = middlewares
}

func setupMiddleware(h http.Handler, c *Core) http.Handler {
	c.Log.Info().Msg("Setup router middlewares")
	return h
}

// Json middleware to set content-type header for json
func jsonMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		h.ServeHTTP(w, r)
	})
}

func loggingMiddleware(h http.Handler, c *Core) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerByte, _ := json.Marshal(r.Header)
		// bodyByte, _ := json.Marshal(r.Body)

		// headerString := string(headerByte)
		c.Log.Debug().
			//Str("agent", r.UserAgent()).
			Str("referer", r.Referer()).
			Str("proto", r.Proto).
			Str("remote_addr", r.RemoteAddr).
			Str("method", r.Method).
			Str("url", r.URL.String()).
			RawJSON("headers", headerByte).
			//RawJSON("body", bodyByte).
			Msg("Request")

		//fmt.Println(headerString)

		h.ServeHTTP(w, r)

	})
}
