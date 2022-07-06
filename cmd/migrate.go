package cmd

import (
	"fmt"

	//   "goat/core"

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
}
