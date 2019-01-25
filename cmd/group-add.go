package cmd

import (
	"fmt"

	"github.com/clockworksoul/cog2/client"
	"github.com/spf13/cobra"
)

const (
	groupAddUse   = "add"
	groupAddShort = "Add a user to an existing group"
	groupAddLong  = "Add a user to an existing group."
)

// GetGroupAddCmd is a command
func GetGroupAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   groupAddUse,
		Short: groupAddShort,
		Long:  groupAddLong,
		RunE:  groupAddCmd,
		Args:  cobra.ExactArgs(2),
	}

	return cmd
}

func groupAddCmd(cmd *cobra.Command, args []string) error {
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
