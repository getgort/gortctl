package cmd

import (
	"fmt"

	"github.com/clockworksoul/cog2/client"
	"github.com/clockworksoul/cog2/data/rest"
	"github.com/spf13/cobra"
)

const (
	bootstrapUse   = "bootstrap"
	bootstrapShort = "Bootstrap a Cog server"
	bootstrapLong  = `Bootstrap can be used on a brand new server, addressable at the specified URL,
to create the administrative user ("admin"). Bootstrapping will only work on
an instance that doesn't yet have any users defined.

Following a successful bootstrapping, the returned password and user
information are added to cogctl's configuration file as a new profile. If this
is the first profile to be added to this configuration file, it will be marked
as the default.

By default, the new profile will be named for the hostname of the server being
bootstrapped. This can be overridden using the -P or --profile flags.`

	bootstrapUsage = `Usage:
  cogctl bootstrap [flags] [URL]

Flags:
  -e, --email string      Email for the bootstrapped user (default "admin@cog")
  -h, --help              help for bootstrap
  -n, --name string       Full name of the bootstrapped user (default "Cog Administrator")
  -p, --password string   Password for the bootstrapped user (default generated)

Global Flags:
  -P, --profile string   Cog profile to use
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

	cmd.Flags().StringVarP(&flagBootstrapEmail, "email", "e", "admin@cog", "Email for the bootstrapped user")
	cmd.Flags().StringVarP(&flagBootstrapName, "name", "n", "Cog Administrator", "Full name of the bootstrapped user")
	cmd.Flags().StringVarP(&flagBootstrapPassword, "password", "p", "", "Password for bootstrapped user (default generated)")

	cmd.SetUsageTemplate(bootstrapUsage)

	return cmd
}

func bootstrapCmd(cmd *cobra.Command, args []string) error {
	entry := client.ProfileEntry{
		Name:      FlagCogProfile,
		URLString: args[0],
	}

	client, err := client.ConnectWithNewProfile(entry)
	if err != nil {
		return err
	}

	user := rest.User{
		Email:    flagBootstrapEmail,
		FullName: flagBootstrapName,
		Password: flagBootstrapPassword,
	}

	// Client Bootstrap will create the cog config if necessary, and append
	// the new credentials to it.
	user, err = client.Bootstrap(user)
	if err != nil {
		return err
	}

	fmt.Printf("User %q created and credentials appended to cog config.\n", user.Username)

	return nil
}
