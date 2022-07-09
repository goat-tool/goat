package core

import (
	"goat/log"

	//"fmt"
	"os"
)

func (c *Core) setupLog(isDebug bool, logFile string) {

	//fmt.Println("Input logfile: ", logFile)
	//fmt.Println("Input isDebug: ", isDebug)

	if logFile == "" {
		logFile = c.conf.Log.File
	}

	if !isDebug {
		//fmt.Println("debug not set by cli -> get it from config")
		isDebug = c.conf.Log.Debug
	}

	log, err := log.New(isDebug, logFile)
	if err != nil {
		log.Error().Err(err).Msg("Setup log error")
		os.Exit(2)
	} else {
		c.Log = log
		c.Log.Info().Msg("Setup conf")
		c.Log.Debug().Str("Host", c.conf.Server.Host).Int("Port", c.conf.Server.Port).Msg("Server")
		c.Log.Debug().Str("File", c.conf.Log.File).Bool("Debug", c.conf.Log.Debug).Msg("Logger")
		c.Log.Debug().Str("Host", c.conf.Database.Host).Str("Username", c.conf.Database.Username).Str("Password", c.conf.Database.Password).Str("Driver", c.conf.Database.Driver).Str("Port", c.conf.Database.Port).Str("Name", c.conf.Database.Name).Msg("Database")
		c.Log.Info().Msg("Setup log")
	}

}
