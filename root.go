package main

import (
	"fmt"
	"os"

	"github.com/clockworksoul/cogctl2/cmd"
	"github.com/spf13/cobra"
)

const (
	rootShort = "Cogctl is a CLI tool for administering a Cog2 chatops server installation"

	rootLong = `Cogctl is a CLI tool for administering a Cog2 chatops server installation.

Did you find a bug or have a suggestion for a new feature? Create an issue at
https://github.com/clockworksoul/cog2/issues.
`
)

var rootCmd = &cobra.Command{
	Use:     "cogctl",
	Short:   rootShort,
	Long:    rootLong,
	Version: Version,
}

func init() {
	rootCmd.AddCommand(cmd.GetBootstrapCmd())
	rootCmd.AddCommand(cmd.GetGroupCmd())
	rootCmd.AddCommand(cmd.GetUserCmd())
	rootCmd.SetVersionTemplate("Cogctl version v" + Version + "\n")
}

// Execute executes
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
