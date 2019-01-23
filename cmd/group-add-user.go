package cmd

import (
	"fmt"

	"github.com/clockworksoul/cog2/client"
	"github.com/spf13/cobra"
)

const (
	groupAddUserUse   = "add-user"
	groupAddUserShort = "Add a user to an existing group"
	groupAddUserLong  = "Add a user to an existing group."
)

// GetGroupAddUserCmd is a command
func GetGroupAddUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   groupAddUserUse,
		Short: groupAddUserShort,
		Long:  groupAddUserLong,
		RunE:  groupAddUserCmd,
		Args:  cobra.ExactArgs(2),
	}

	return cmd
}

func groupAddUserCmd(cmd *cobra.Command, args []string) error {
	groupname := args[0]
	username := args[1]

	client, err := client.Connect(FlagCogProfile)
	if err != nil {
		return printError(err)
	}

	err = client.GroupMemberAdd(groupname, username)
	if err != nil {
		return printError(err)
	}

	fmt.Printf("User added to %s: %s\n", groupname, username)

	return nil
}
