package cmd

import (
	"fmt"

	"github.com/clockworksoul/cog2/client"
	"github.com/spf13/cobra"
)

const (
	userDeleteUse   = "delete"
	userDeleteShort = "Delete an existing user"
	userDeleteLong  = "Delete an existing user."
)

// GetUserDeleteCmd is a command
func GetUserDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   userDeleteUse,
		Short: userDeleteShort,
		Long:  userDeleteLong,
		RunE:  userDeleteCmd,
		Args:  cobra.ExactArgs(1),
	}

	return cmd
}

func userDeleteCmd(cmd *cobra.Command, args []string) error {
	client, err := client.Connect(FlagCogProfile)
	if err != nil {
		return printError(err)
	}

	username := args[0]

	user, err := client.UserGet(username)
	if err != nil {
		return printError(err)
	}

	fmt.Printf("Deleting user %s (%s)\n", user.Username, user.Password)

	err = client.UserDelete(user.Username)
	if err != nil {
		return printError(err)
	}

	fmt.Println("Delete successful")

	return nil
}
