package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"

	"github.com/clockworksoul/gort/client"
	gorterr "github.com/clockworksoul/gort/errors"
	"github.com/clockworksoul/gortctl/cmd"
	"github.com/spf13/cobra"
)

const (
	rootShort = "Gortctl is a CLI tool for administering a gort chatops server installation"

	rootLong = `Manage Gort via its REST API on the command line.

Did you find a bug or have a suggestion for a new feature? Create an issue at
https://github.com/clockworksoul/gort/issues.
`
)

var rootCmd = &cobra.Command{
	Use:                        "gortctl",
	Short:                      rootShort,
	Long:                       rootLong,
	Version:                    Version,
	SilenceUsage:               true,
	SilenceErrors:              true,
	SuggestionsMinimumDistance: 2,
}

func init() {
	rootCmd.AddCommand(cmd.GetBootstrapCmd())
	rootCmd.AddCommand(cmd.GetBundleCmd())
	rootCmd.AddCommand(cmd.GetGroupCmd())
	rootCmd.AddCommand(cmd.GetUserCmd())
	rootCmd.SetVersionTemplate("Gortctl version v" + Version + "\n")

	rootCmd.PersistentFlags().StringVarP(&cmd.FlagGortProfile, "profile", "P", "", "Gort profile to use")
}

// Execute executes
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		printErr(err)
		os.Exit(1)
	}
}

func printErr(err error) error {
	var msg string

	switch {
	case reflect.TypeOf(err) == reflect.TypeOf(gorterr.NestedError{}):
		nerr := err.(gorterr.NestedError)
		fmt.Fprintln(os.Stderr, "Error:", nerr.Message)

		if true { // TODO Add "verbose" flag
			suberr := nerr.Err

			for reflect.TypeOf(suberr) == reflect.TypeOf(gorterr.NestedError{}) {
				nerr = suberr.(gorterr.NestedError)
				suberr = nerr.Err

				fmt.Fprintf(os.Stderr, "Caused by: %s\n", nerr.Message)
			}

			fmt.Fprintf(os.Stderr, "Caused by: %s\n", suberr.Error())
		}
	case reflect.TypeOf(err) == reflect.TypeOf(client.Error{}):
		cerr := err.(client.Error)
		msg = getClientErrorMessage(cerr)
		fmt.Fprintln(os.Stderr, "Error:", msg)
	default:
		msg = err.Error()
		fmt.Fprintln(os.Stderr, "Error:", msg)
	}

	return err
}

func getClientErrorMessage(cerr client.Error) string {
	status := cerr.Status()

	switch status {
	case 0:
		return fmt.Sprintf("Could not establish HTTP connection to %s. "+
			"Please check your host, user, and password settings.",
			cerr.Profile().URLString)

	case http.StatusNoContent:
		return "No items found"

	default:
		return fmt.Sprintf("%d %s", status, cerr.Error())
	}
}
