package cmd

import (
	"fmt"
	"strings"

	"github.com/clockworksoul/gort/client"
	"github.com/spf13/cobra"
)

const (
	groupInfoUse   = "info"
	groupInfoShort = "Retrieve information about an existing group"
	groupInfoLong  = "Retrieve information about an existing group."
)

// GetGroupInfoCmd is a command
func GetGroupInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   groupInfoUse,
		Short: groupInfoShort,
		Long:  groupInfoLong,
		RunE:  groupInfoCmd,
		Args:  cobra.ExactArgs(1),
	}

	return cmd
}

func checkArgs(cmd *cobra.Command, args []string, argf cobra.PositionalArgs) error {
	var err error

	if args != nil {
		err = argf(cmd, args)
	}

	return err
}

func groupInfoCmd(cmd *cobra.Command, args []string) error {
	groupname := args[0]

	gortClient, err := client.Connect(FlagGortProfile)
	if err != nil {
		return err
	}

	//
	// TODO Maybe multiplex the following queries with gofuncs?
	// (when there's more than one)
	//

	users, err := gortClient.GroupMemberList(groupname)
	if err != nil {
		return err
	}

	// TODO Add "roles" here when its supported.

	const format = `Name   %s
Users  %s
`

	fmt.Printf(format, groupname, strings.Join(userNames(users), ", "))

	return nil
}
