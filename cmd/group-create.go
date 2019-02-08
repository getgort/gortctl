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

	c, err := client.Connect(FlagCogProfile)
	if err != nil {
		return err
	}

	// Only allow this operation if the group doesn't already exist.
	exists, err := c.GroupExists(groupname)
	if err != nil {
		return err
	}
	if exists {
		return client.ErrResourceExists
	}

	group := rest.Group{Name: groupname}

	// Client GroupCreate will create the cog config if necessary, and append
	// the new credentials to it.
	err = c.GroupSave(group)
	if err != nil {
		return err
	}

	fmt.Printf("Group %q created.\n", group.Name)

	return nil
}
