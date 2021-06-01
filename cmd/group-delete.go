package cmd

import (
	"fmt"

	"github.com/getgort/gort/client"
	"github.com/spf13/cobra"
)

const (
	groupDeleteUse   = "delete"
	groupDeleteShort = "Delete an existing group"
	groupDeleteLong  = "Delete an existing group."
)

// GetGroupDeleteCmd is a command
func GetGroupDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   groupDeleteUse,
		Short: groupDeleteShort,
		Long:  groupDeleteLong,
		RunE:  groupDeleteCmd,
		Args:  cobra.ExactArgs(1),
	}

	return cmd
}

func groupDeleteCmd(cmd *cobra.Command, args []string) error {
	gortClient, err := client.Connect(FlagGortProfile)
	if err != nil {
		return err
	}

	groupname := args[0]

	group, err := gortClient.GroupGet(groupname)
	if err != nil {
		return err
	}

	fmt.Printf("Deleting group %s... ", group.Name)

	err = gortClient.GroupDelete(group.Name)
	if err != nil {
		return err
	}

	fmt.Println("Successful")

	return nil
}
