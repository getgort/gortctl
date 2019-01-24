package cmd

import (
	"fmt"

	"github.com/clockworksoul/cog2/client"
	"github.com/clockworksoul/cog2/data/rest"
	"github.com/spf13/cobra"
)

const (
	groupCreateUse   = "create"
	groupCreateShort = "Create an existing group"
	groupCreateLong  = "Create an existing group."
)

// GetGroupCreateCmd is a command
func GetGroupCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   groupCreateUse,
		Short: groupCreateShort,
		Long:  groupCreateLong,
		RunE:  groupCreateCmd,
		Args:  cobra.ExactArgs(1),
	}

	return cmd
}

func groupCreateCmd(cmd *cobra.Command, args []string) error {
	groupname := args[0]

	client, err := client.Connect(FlagCogProfile)
	if err != nil {
		return printError(err)
	}

	// Only allow this operation if the group doesn't already exist.
	exists, err := client.GroupExists(groupname)
	if err != nil {
		return printError(err)
	}

	if exists {
		return printErrorf("Group %s already exists.\n", groupname)
	}

	group := rest.Group{Name: groupname}

	// Client GroupCreate will create the cog config if necessary, and append
	// the new credentials to it.
	err = client.GroupSave(group)
	if err != nil {
		return printError(err)
	}

	fmt.Printf("Group %q created.\n", group.Name)

	return nil
}
