package cmd

import (
	"github.com/spf13/cobra"
)

const (
	userUse   = "user"
	userShort = "Perform operations on users"
	userLong  = "Allows you to perform user administration."
)

// # gortctl user --help
// Usage: gortctl user [OPTIONS] COMMAND [ARGS]...
//
//   Manage Gort users.
//
//   If invoked without a subcommand, lists all the users on the server.
//
// Options:
//   --help  Show this message and exit.
//
// Commands:
//   create                  Create a new user.
//   delete                  Deletes a user.
//   info                    Get info about a specific user by username.
//   password-reset          Reset user password with a token.
//   password-reset-request  Request a password reset.
//   update                  Updates an existing user.

// GetUserCmd user
func GetUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   userUse,
		Short: userShort,
		Long:  userLong,
	}

	cmd.AddCommand(GetUserCreateCmd())
	cmd.AddCommand(GetUserDeleteCmd())
	cmd.AddCommand(GetUserInfoCmd())
	cmd.AddCommand(GetUserListCmd())
	cmd.AddCommand(GetUserUpdateCmd())

	return cmd
}
