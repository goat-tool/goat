package cmd

import (
	"fmt"
	"os"

	"goat/core"
	"goat/services/test"
	"goat/services/user"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate database",
	Long:  `Creates database structure based on models`,
	Run:   migrate,
}

func migrate(cmd *cobra.Command, args []string) {
	fmt.Println("Todo: migrate command")

	appCore, err := core.New(cfgFile, isDebug, logFile)
	if err != nil {
		//Log.Debug().Err(err).Msg("initCore error")
		os.Exit(2)
	}

	appCore.NewDatabase()
	appCore.Database.AutoMigrate(
		test.Test{},
		user.User{},
	)

}
