package cmd

import (
	"github.com/spf13/cobra"
)

const (
	groupUse   = "group"
	groupShort = "Perform operations on groups"
	groupLong  = "Allows you to create or delete groups."
)

// GetGroupCmd group
func GetGroupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   groupUse,
		Short: groupShort,
		Long:  groupLong,
	}

	return cmd
}
