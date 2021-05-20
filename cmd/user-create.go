package cmd

import (
	"fmt"

	"github.com/clockworksoul/gort/client"
	"github.com/clockworksoul/gort/data/rest"
	"github.com/spf13/cobra"
)

const (
	userCreateUse   = "create"
	userCreateShort = "Create a new user"
	userCreateLong  = "Create a new user."
)

var (
	flagUserCreateEmail    string
	flagUserCreateName     string
	flagUserCreatePassword string
	flagUserCreateProfile  string
)

// GetUserCreateCmd is a command
func GetUserCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   userCreateUse,
		Short: userCreateShort,
		Long:  userCreateLong,
		RunE:  userCreateCmd,
		Args:  cobra.ExactArgs(1),
	}

	cmd.Flags().StringVarP(&flagUserCreateEmail, "email", "e", "", "Email for the user (required)")
	cmd.Flags().StringVarP(&flagUserCreateName, "name", "n", "", "Full name of the user (required)")
	cmd.Flags().StringVarP(&flagUserCreatePassword, "password", "p", "", "Password for user (required)")

	cmd.MarkFlagRequired("email")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("password")

	return cmd
}

func userCreateCmd(cmd *cobra.Command, args []string) error {
	username := args[0]

	c, err := client.Connect(FlagGortProfile)
	if err != nil {
		return err
	}

	// Only allow this operation if the user doesn't already exist.
	exists, err := c.UserExists(username)
	if err != nil {
		return err
	}
	if exists {
		return client.ErrResourceExists
	}

	user := rest.User{
		Email:    flagUserCreateEmail,
		FullName: flagUserCreateName,
		Password: flagUserCreatePassword,
		Username: username,
	}

	err = c.UserSave(user)
	if err != nil {
		return err
	}

	fmt.Printf("User %q created.\n", user.Username)

	return nil
}
