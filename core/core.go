package core

import (
	"goat/conf"
	"goat/log"

	"goat/api"
	"goat/services"
	"goat/services/health"

	"context"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type Core struct {
	Log  *log.Logger
	conf *conf.Config
	//	Database *mongo.Client
	httpServer *http.Server
	router     *mux.Router
	handler    http.Handler
	api        *api.Api
	services   *services.Services
	state      health.State
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

	c.state = health.StateStarting
	c.router = mux.NewRouter()

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
	c.setupDatabase()
	c.setupServices()
	c.setupApi()
	c.setupRouter()

	c.Log.Info().Msg("New core done")
	c.state = health.StateRunning
	return c, nil
}

func (c *Core) StartServer() {

	var err error
	c.Log.Info().Msg("Starting server")

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
		if c.state == health.StateStopping || c.state == health.StateStopped {
			c.Log.Debug().Msgf("Skipping restarts of server because app is not in running state. State is:", c.state)
			return
		}
		if err = listen(); err != nil {
			time.Sleep(2 * time.Second)
			c.Log.Error().Err(err).Msgf("Error on", c.httpServer.Addr)
			c.Log.Debug().Msgf("Restarting server", c.httpServer.Addr)
			continue
		}
		return
	}
}

func (c *Core) Shutdown(ctx context.Context) error {

	defer func() {
		//fmt.Println("defer...........................")
	}()

	c.context.cancel()
	done := make(chan struct{})
	go func() {
		c.wg.Wait()
		close(done)
	}()

	// loop tru shutdownfuncs
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
