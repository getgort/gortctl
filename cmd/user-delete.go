package cmd

import (
	"fmt"

	"github.com/clockworksoul/gort/client"
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
	gortClient, err := client.Connect(FlagGortProfile)
	if err != nil {
		return err
	}

	username := args[0]

	user, err := gortClient.UserGet(username)
	if err != nil {
		return err
	}

	fmt.Printf("Deleting user %s (%s)... ", user.Username, user.Email)

	err = gortClient.UserDelete(user.Username)
	if err != nil {
		return err
	}

	fmt.Println("Successful.")

	return nil
}
