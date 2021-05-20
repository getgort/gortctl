package cmd

import (
	"fmt"

	"github.com/clockworksoul/gort/client"
	"github.com/spf13/cobra"
)

const (
	groupRemoveUse   = "add"
	groupRemoveShort = "Remove a user from an existing group"
	groupRemoveLong  = "Remove a user from an existing group."
)

// GetGroupRemoveCmd is a command
func GetGroupRemoveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   groupRemoveUse,
		Short: groupRemoveShort,
		Long:  groupRemoveLong,
		RunE:  groupRemoveCmd,
		Args:  cobra.ExactArgs(2),
	}

	return cmd
}

func groupRemoveCmd(cmd *cobra.Command, args []string) error {
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
