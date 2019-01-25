package cmd

import (
	"github.com/spf13/cobra"
)

const (
	groupUse   = "group"
	groupShort = "Perform operations on groups"
	groupLong  = "Allows you to perform group administration."
)

// # cogctl group --help
// Usage: cogctl group [OPTIONS] COMMAND [ARGS]...
//
//   Manage Cog user groups.
//
//   If invoked without a subcommand, lists all user groups.
//
// Options:
//   --help  Show this message and exit.
//
// Commands:
//   add     Add one or more users to a group.
//   create  Create a new user group.
//   delete  Delete a group.
//   grant   Grant one or more roles to a group.
//   info    Show info on a specific group.
//   remove  Remove one or more users from a group.
//   rename  Rename a user group.
//   revoke  Revoke one or more roles from a group.

// GetGroupCmd group
func GetGroupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   groupUse,
		Short: groupShort,
		Long:  groupLong,
	}

	cmd.AddCommand(GetGroupAddCmd())
	cmd.AddCommand(GetGroupCreateCmd())
	cmd.AddCommand(GetGroupDeleteCmd())
	cmd.AddCommand(GetGroupInfoCmd())
	cmd.AddCommand(GetGroupListCmd())

	return cmd
}
