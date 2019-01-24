package cmd

import (
	"github.com/spf13/cobra"
)

const (
	userUse   = "user"
	userShort = "Perform operations on users"
	userLong  = "Allows you to perform user administration."
)

// GetUserCmd user
func GetUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   userUse,
		Short: userShort,
		Long:  userLong,
	}

	cmd.AddCommand(GetUserDeleteCmd())
	cmd.AddCommand(GetUserInfoCmd())
	cmd.AddCommand(GetUserListCmd())

	return cmd
}
