package cmd

import (
	"github.com/spf13/cobra"
)

const (
	groupUse   = "group"
	groupShort = "Perform operations on groups"
	groupLong  = "Allows you to perform group administration."
)

// GetGroupCmd group
func GetGroupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   groupUse,
		Short: groupShort,
		Long:  groupLong,
	}

	cmd.AddCommand(GetGroupDeleteCmd())
	cmd.AddCommand(GetGroupInfoCmd())
	cmd.AddCommand(GetGroupListCmd())

	return cmd
}
