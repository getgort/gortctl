package cmd

import (
	"fmt"

	"github.com/clockworksoul/gort/client"
	"github.com/spf13/cobra"
)

const (
	groupRemoveUserUse   = "remove-user"
	groupRemoveUserShort = "Remove a user from an existing group"
	groupRemoveUserLong  = "Remove a user from an existing group."
)

// GetGroupRemoveUserCmd is a command
func GetGroupRemoveUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   groupRemoveUserUse,
		Short: groupRemoveUserShort,
		Long:  groupRemoveUserLong,
		RunE:  groupRemoveUserCmd,
		Args:  cobra.ExactArgs(2),
	}

	return cmd
}

func groupRemoveUserCmd(cmd *cobra.Command, args []string) error {
	groupname := args[0]
	username := args[1]

	gortClient, err := client.Connect(FlagGortProfile)
	if err != nil {
		return err
	}

	err = gortClient.GroupMemberDelete(groupname, username)
	if err != nil {
		return err
	}

	fmt.Printf("User removed from %s: %s\n", groupname, username)

	return nil
}
