package cmd

import (
	"fmt"

	"github.com/clockworksoul/cog2/client"
	"github.com/spf13/cobra"
)

const (
	userListUse   = "list"
	userListShort = "List all existing users"
	userListLong  = "List all existing users."
)

// GetUserListCmd is a command
func GetUserListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   userListUse,
		Short: userListShort,
		Long:  userListLong,
		RunE:  userListCmd,
	}

	return cmd
}

func userListCmd(cmd *cobra.Command, args []string) error {
	const format = "%-10s%-20s%s\n"

	cogClient, err := client.Connect(FlagCogProfile)
	if err != nil {
		return err
	}

	users, err := cogClient.UserList()
	if err != nil {
		return err
	}

	fmt.Printf(format, "USERNAME", "FULL NAME", "EMAIL ADDRESS")
	for _, u := range users {
		fmt.Printf(format, u.Username, u.FullName, u.Email)
	}

	return nil
}
