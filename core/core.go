package core

import (
	"goat/conf"
	"goat/log"

	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Core struct {
	// log holds a pointer to the logger
	Log *log.Logger

	// config holds a pointer to the loaded configuration
	conf *conf.Config

	// server holds a pointer to the http server
	httpServer *http.Server

	// router is the HTTP mux
	//router *mux.Router

	// handler is the http handler from the mux Router for the http server
	handler http.Handler

	// api holds the api handler for the rest api
	//api *api.Api

	// services holds the internal Services for global usage
	//services *services.Services

	// state is set to the current state of the application
	//state health.State

	// wg holds registered processes for graceful shutdown
	wg *sync.WaitGroup

	// context holds global context
	context globalContext

	// shutdownFuncs runs at shutdown
	shutdownFuncs []func() error
}

type globalContext struct {
	cancel context.CancelFunc
	ctx    context.Context
}

// New Core
func New(cfgFile string, isDebug bool, logFile string) (*Core, error) {
	c := &Core{}

	ctx, cancel := context.WithCancel(context.Background())
	c.wg = &sync.WaitGroup{}
	c.context = globalContext{
		cancel: cancel,
		ctx:    ctx,
	}

	c.setupConf(cfgFile)
	c.setupLog(isDebug, logFile)
	// c.setupTranslator()
	// c.setupValidator()
	// c.setupDatabase()
	c.setupServices()
	c.setupApi()
	c.setupRouter()

	//c.router = mux.NewRouter()

	c.Log.Info().Msg("Init core done")

	return c, nil
}

func (c *Core) StartServer() {

	var err error
	c.Log.Debug().Msg("Starting server")

	c.httpServer = &http.Server{
		Addr:         c.conf.Server.URL(),
		Handler:      c.handler,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	listen := func() error {
		if err = c.httpServer.ListenAndServe(); err != http.ErrServerClosed {
			c.Log.Error().Err(err).Msg("failed to start http server")
			return err
		}

		return nil
	}

	c.Log.Info().Str("URL", c.httpServer.Addr).Msg("Server listen")

	for {
		// if c.state == health.StateStopping || c.state == health.StateStopped {
		// 	logger.Tracef("skipping restarts of server because app is not in running state: state is %d", c.state)
		// 	return
		// }
		if err = listen(); err != nil {
			time.Sleep(2 * time.Second)
			fmt.Println("restarting")
			continue
		}
		return
	}

	// if err = listen(); err != nil {
	// 	c.log.Error().Err(err).Msgf("Error on: ", c.httpServer.Addr)
	// }

}

func (c *Core) Shutdown(ctx context.Context) error {

	defer func() {
		//fmt.Println("defer...........................")
	}()

	// if c.httpServer != nil {
	// 	if err := c.httpServer.Shutdown(ctx); err != nil {
	// 		c.Log.Error().Msgf("server shutdown error", err)
	// 	} else {
	// 		c.Log.Debug().Msg("http server stopped")
	// 	}
	// }

	c.context.cancel()
	done := make(chan struct{})
	go func() {
		c.wg.Wait()
		close(done)
	}()

	for _, fn := range c.shutdownFuncs {
		funcErr := fn()
		if funcErr != nil {
			c.Log.Error().Msgf("Shutdown func failed with error", funcErr)
		}
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-done:
	}

	return nil
}

// // registerShutdownFunc is used to register a new shutdown function
// func (c *Core) registerShutdownFunc(fn func() error) {
// 	c.shutdownFuncs = append(c.shutdownFuncs, fn)
// }
