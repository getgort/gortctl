package cmd

import (
	"fmt"
	"strings"

	"github.com/clockworksoul/cog2/client"
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

func groupInfoCmd(cmd *cobra.Command, args []string) error {
	groupname := args[0]

	client, err := client.Connect(FlagCogProfile)
	if err != nil {
		return printError(err)
	}

	//
	// TODO Maybe multiplex the following queries with gofuncs?
	// (when there's more than one)
	//

	users, err := client.GroupMemberList(groupname)
	if err != nil {
		return printError(err)
	}

	// TODO Add "roles" here when its supported.

	const format = `Name   %s
Users  %s
`

	fmt.Printf(format, groupname, strings.Join(userNames(users), ""))

	return nil
}
