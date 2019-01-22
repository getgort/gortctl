package cmd

import (
	"github.com/spf13/cobra"
)

const (
	userUse   = "user"
	userShort = "Perform operations on users"
	userLong  = "Allows you to create or delete users."
)

// GetUserCmd user
func GetUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   userUse,
		Short: userShort,
		Long:  userLong,
	}

	return cmd
}
