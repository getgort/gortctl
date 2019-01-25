package cmd

import (
	"fmt"

	"github.com/clockworksoul/cog2/client"
	"github.com/clockworksoul/cog2/data/rest"
	"github.com/spf13/cobra"
)

const (
	userUpdateUse   = "update"
	userUpdateShort = "Update an existing user"
	userUpdateLong  = "Update an existing user."
)

var (
	flagUserUpdateEmail    string
	flagUserUpdateName     string
	flagUserUpdatePassword string
	flagUserUpdateProfile  string
)

// GetUserUpdateCmd is a command
func GetUserUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   userUpdateUse,
		Short: userUpdateShort,
		Long:  userUpdateLong,
		RunE:  userUpdateCmd,
		Args:  cobra.ExactArgs(1),
	}

	cmd.Flags().StringVarP(&flagUserUpdateEmail, "email", "e", "", "Email for the user")
	cmd.Flags().StringVarP(&flagUserUpdateName, "name", "n", "", "Full name of the user")
	cmd.Flags().StringVarP(&flagUserUpdatePassword, "password", "p", "", "Password for user")

	return cmd
}

func userUpdateCmd(cmd *cobra.Command, args []string) error {
	username := args[0]

	client, err := client.Connect(FlagCogProfile)
	if err != nil {
		return printError(err)
	}

	// Only allow this operation if the user already exists.
	exists, err := client.UserExists(username)
	if err != nil {
		return printError(err)
	}

	if !exists {
		return printErrorf("User %q doesn't exist.\n", username)
	}

	// Empty fields will not be overwritten.
	user := rest.User{
		Email:    flagUserUpdateEmail,
		FullName: flagUserUpdateName,
		Password: flagUserUpdatePassword,
		Username: username,
	}

	err = client.UserSave(user)
	if err != nil {
		return printError(err)
	}

	fmt.Printf("User %q updated.\n", user.Username)

	return nil
}
