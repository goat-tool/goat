package core

import (
	"goat/conf"

	"fmt"
	"os"
)

func (c *Core) setupConf(cfgFile string) {

	//fmt.Println("input cfgfile: ", cfgFile)

	conf, err := conf.New(cfgFile)
	if err != nil {
		fmt.Println("InitConfig error")
		os.Exit(2)
	} else {
		c.conf = conf
	}

}
