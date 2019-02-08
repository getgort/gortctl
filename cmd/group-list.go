package cmd

import (
	"fmt"

	"github.com/clockworksoul/cog2/client"
	"github.com/spf13/cobra"
)

const (
	groupListUse   = "list"
	groupListShort = "List all existing groups"
	groupListLong  = "List all existing groups."
)

// GetGroupListCmd is a command
func GetGroupListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   groupListUse,
		Short: groupListShort,
		Long:  groupListLong,
		RunE:  groupListCmd,
	}

	return cmd
}

func groupListCmd(cmd *cobra.Command, args []string) error {
	const format = "%s\n"

	client, err := client.Connect(FlagCogProfile)
	if err != nil {
		return err
	}

	groups, err := client.GroupList()
	if err != nil {
		return err
	}

	fmt.Printf(format, "GROUP NAME")
	for _, g := range groups {
		fmt.Printf(format, g.Name)
	}

	return nil
}
