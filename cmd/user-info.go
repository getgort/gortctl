package cmd

import (
	"fmt"
	"strings"

	"github.com/clockworksoul/gort/client"
	"github.com/spf13/cobra"
)

const (
	userInfoUse   = "info"
	userInfoShort = "Retrieve information about an existing user"
	userInfoLong  = "Retrieve information about an existing user."
)

// GetUserInfoCmd is a command
func GetUserInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   userInfoUse,
		Short: userInfoShort,
		Long:  userInfoLong,
		RunE:  userInfoCmd,
		Args:  cobra.ExactArgs(1),
	}

	return cmd
}

func userInfoCmd(cmd *cobra.Command, args []string) error {
	gortClient, err := client.Connect(FlagGortProfile)
	if err != nil {
		return err
	}

	//
	// TODO Maybe multiplex the following queries with gofuncs?
	//

	username := args[0]

	user, err := gortClient.UserGet(username)
	if err != nil {
		return err
	}

	groups, err := gortClient.UserGroupList(username)
	if err != nil {
		return err
	}

	const format = `Name       %s
Full Name  %s
Email      %s
Groups     %s
`

	fmt.Printf(format, user.Username, user.FullName, user.Email, strings.Join(groupNames(groups), ", "))

	return nil
}
