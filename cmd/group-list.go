package cmd

import (
	"fmt"

	"github.com/clockworksoul/gort/client"
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

	gortClient, err := client.Connect(FlagGortProfile)
	if err != nil {
		return err
	}

	groups, err := gortClient.GroupList()
	if err != nil {
		return err
	}

	fmt.Printf(format, "GROUP NAME")
	for _, g := range groups {
		fmt.Printf(format, g.Name)
	}

	return nil
}
