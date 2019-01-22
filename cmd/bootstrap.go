package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/clockworksoul/cog2/client"
	"github.com/clockworksoul/cog2/data/rest"
	"github.com/spf13/cobra"
)

const (
	bootstrapUse   = "bootstrap"
	bootstrapShort = "Bootstrap a new server."
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

	cmd.Flags().StringVarP(&flagBootstrapEmail, "email", "e", "", "Email for the bootstrapped user")
	cmd.Flags().StringVarP(&flagBootstrapName, "name", "n", "", "Full name of the bootstrapped user")
	cmd.Flags().StringVarP(&flagBootstrapPassword, "password", "p", "", "Password for bootstrapped user")
	cmd.Flags().StringVarP(&flagBootstrapProfile, "profile", "r", "", "Name to use for the server profile")
	cmd.Flags().StringVarP(&flagBootstrapUser, "user", "u", "", "Username for the bootstrapped user")

	return cmd
}

func bootstrapCmd(cmd *cobra.Command, args []string) error {
	client, err := client.Connect(args[0])
	if err != nil {
		log.Fatal(err)
	}

	// Does the user want a specific email?
	// If not, Cog will apply the default ("cog@localhost") and return it
	email := flagBootstrapEmail

	// Does the user want a specific password?
	// If not, Cog will generate one automatically and return it
	password := flagBootstrapPassword

	// Does the user want a specific name?
	// If not, Cog will generate one automatically and return it
	fullName := flagBootstrapName

	// The default profile name is the hostname of the Cog server
	profile := flagBootstrapProfile
	if profile == "" {
		profile = client.Host().Hostname()
	}

	// Does the user want a specific username?
	// If not, Cog will apply the default ("admin") and return it
	username := flagBootstrapUser

	user := rest.User{
		Email:    email,
		FullName: fullName,
		Password: password,
		Username: username,
	}

	user, err = client.Bootstrap(user)
	if err != nil {
		log.Fatal(err)
	}

	bytes, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println(string(bytes))

	return nil
}
