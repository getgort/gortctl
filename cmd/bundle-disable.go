package cmd

import (
	"fmt"

	"github.com/getgort/gort/client"
	"github.com/spf13/cobra"
)

const (
	bundleDisableUse   = "disable"
	bundleDisableShort = "Disable a bundle"
	bundleDisableLong  = "Disable a bundle."
)

// cogctl bundle disable --help
// Usage: cogctl bundle disable [OPTIONS] NAME

//   Disable a bundle by name.

// Options:
//   --help  Show this message and exit.

// GetBundleDisableCmd is a command
func GetBundleDisableCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   bundleDisableUse,
		Short: bundleDisableShort,
		Long:  bundleDisableLong,
		RunE:  bundleDisableCmd,
		Args:  cobra.ExactArgs(2),
	}

	return cmd
}

func bundleDisableCmd(cmd *cobra.Command, args []string) error {
	bundleName := args[0]
	bundleVersion := args[1]

	c, err := client.Connect(FlagGortProfile)
	if err != nil {
		return err
	}

	err = c.BundleDisable(bundleName, bundleVersion)
	if err != nil {
		return err
	}

	fmt.Printf("Bundle \"%s\" disabled.\n", bundleName)

	return nil
}
