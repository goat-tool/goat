package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of goat",
	Long:  `All software has versions. This is goat's`,
	Run:   version,
}

func version(cmd *cobra.Command, args []string) {
	fmt.Println("goat version 0.0.1")
}
