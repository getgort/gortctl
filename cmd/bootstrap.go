package cmd

import (
	"fmt"

	"github.com/clockworksoul/gort/client"
	"github.com/clockworksoul/gort/data/rest"
	"github.com/spf13/cobra"
)

const (
	bootstrapUse   = "bootstrap"
	bootstrapShort = "Bootstrap a Gort server"
	bootstrapLong  = `Bootstrap can be used on a brand new server, addressable at the specified URL,
to create the administrative user ("admin"). Bootstrapping will only work on
an instance that doesn't yet have any users defined.

Following a successful bootstrapping, the returned password and user
information are added to gortctl's configuration file as a new profile. If this
is the first profile to be added to this configuration file, it will be marked
as the default.

By default, the new profile will be named for the hostname of the server being
bootstrapped. This can be overridden using the -P or --profile flags.`

	bootstrapUsage = `Usage:
  gortctl bootstrap [flags] [URL]

Flags:
  -e, --email string      Email for the bootstrapped user (default "admin@gort")
  -h, --help              help for bootstrap
  -n, --name string       Full name of the bootstrapped user (default "Gort Administrator")
  -p, --password string   Password for the bootstrapped user (default generated)

Global Flags:
  -P, --profile string   Gort profile to use
`
)

var (
	flagBootstrapEmail    string
	flagBootstrapName     string
	flagBootstrapPassword string
)

// GetBootstrapCmd bootstrap
func GetBootstrapCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   bootstrapUse,
		Short: bootstrapShort,
		Long:  bootstrapLong,
		RunE:  bootstrapCmd,
		Args:  cobra.ExactArgs(1),
	}

	cmd.Flags().StringVarP(&flagBootstrapEmail, "email", "e", "admin@gort", "Email for the bootstrapped user")
	cmd.Flags().StringVarP(&flagBootstrapName, "name", "n", "Gort Administrator", "Full name of the bootstrapped user")
	cmd.Flags().StringVarP(&flagBootstrapPassword, "password", "p", "", "Password for bootstrapped user (default generated)")

	cmd.SetUsageTemplate(bootstrapUsage)

	return cmd
}

func bootstrapCmd(cmd *cobra.Command, args []string) error {
	entry := client.ProfileEntry{
		Name:      FlagGortProfile,
		URLString: args[0],
	}

	gortClient, err := client.ConnectWithNewProfile(entry)
	if err != nil {
		return err
	}

	user := rest.User{
		Email:    flagBootstrapEmail,
		FullName: flagBootstrapName,
		Password: flagBootstrapPassword,
	}

	// Client Bootstrap will create the gort config if necessary, and append
	// the new credentials to it.
	user, err = gortClient.Bootstrap(user)
	if err != nil {
		return err
	}

	fmt.Printf("User %q created and credentials appended to gort config.\n", user.Username)

	return nil
}
