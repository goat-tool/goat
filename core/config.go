package core

import (
	"goat/conf"

	"fmt"
	"os"
)

func (c *Core) newConf(cfgFile string) {
	conf, err := conf.New(cfgFile)
	if err != nil {
		fmt.Println("Setup config error")
		os.Exit(2)
	} else {
		c.Conf = conf
	}
}
