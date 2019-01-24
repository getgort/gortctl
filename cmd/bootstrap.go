package cmd

import (
	"fmt"

	"github.com/clockworksoul/cog2/client"
	"github.com/clockworksoul/cog2/data/rest"
	"github.com/spf13/cobra"
)

const (
	bootstrapUse   = "bootstrap"
	bootstrapShort = "Bootstrap a Cog serverr"
	bootstrapLong  = `Bootstrap can be used on a brand new server to create the administrative user (admin).
	
It will only work on an instance that doesn't yet have any users defined.`
)

var (
	flagBootstrapEmail    string
	flagBootstrapName     string
	flagBootstrapPassword string
	flagBootstrapProfile  string
	flagBootstrapUser     string
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
	cmd.Flags().StringVarP(&flagBootstrapPassword, "password", "p", "", "Password for bootstrapped user")
	cmd.Flags().StringVarP(&flagBootstrapUser, "username", "u", "admin", "Username for the bootstrapped user")

	return cmd
}

func bootstrapCmd(cmd *cobra.Command, args []string) error {
	entry := client.ProfileEntry{
		Name:      flagBootstrapProfile,
		URLString: args[0],
	}

	client, err := client.ConnectWithNewProfile(entry)
	if err != nil {
		return printError(err)
	}

	user := rest.User{
		Email:    flagBootstrapEmail,
		FullName: flagBootstrapName,
		Password: flagBootstrapPassword,
		Username: flagBootstrapUser,
	}

	// Client Bootstrap will create the cog config if necessary, and append
	// the new credentials to it.
	user, err = client.Bootstrap(user)
	if err != nil {
		return printError(err)
	}

	fmt.Printf("User %q created and credentials appended to cog config.\n", user.Username)

	return nil
}
