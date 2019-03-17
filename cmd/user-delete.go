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
	cogClient, err := client.Connect(FlagCogProfile)
	if err != nil {
		return err
	}

	username := args[0]

	user, err := cogClient.UserGet(username)
	if err != nil {
		return err
	}

	fmt.Printf("Deleting user %s (%s)... ", user.Username, user.Email)

	err = cogClient.UserDelete(user.Username)
	if err != nil {
		return err
	}

	fmt.Println("Successful.")

	return nil
}
