package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	cfgFile string
	isDebug bool
	logFile string

	rootCmd = &cobra.Command{
		Use:   "goat",
		Short: "Goat is a go api template",
		Long: `Goat is a go api template empowering new applications.
This application is under heavy development.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	//cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Config file (default is $HOME/.goat/config.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&isDebug, "debug", "d", false, "Set log level to debug")
	rootCmd.PersistentFlags().StringVarP(&logFile, "log", "l", "", "Set log file (default is $HOME/.goat/goat.log)")
}
