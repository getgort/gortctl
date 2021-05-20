package cmd

import (
	"fmt"

	"github.com/clockworksoul/gort/client"
	"github.com/spf13/cobra"
)

const (
	bundleEnableUse   = "enable"
	bundleEnableShort = "Enable a bundle"
	bundleEnableLong  = "Enable a bundle."
)

// GetBundleEnableCmd is a command
func GetBundleEnableCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   bundleEnableUse,
		Short: bundleEnableShort,
		Long:  bundleEnableLong,
		RunE:  bundleEnableCmd,
		Args:  cobra.ExactArgs(2),
	}

	return cmd
}

func bundleEnableCmd(cmd *cobra.Command, args []string) error {
	bundleName := args[0]
	bundleVersion := args[1]

	c, err := client.Connect(FlagGortProfile)
	if err != nil {
		return err
	}

	err = c.BundleEnable(bundleName, bundleVersion)
	if err != nil {
		return err
	}

	fmt.Printf("Bundle %s v%s version enabled.\n", bundleName, bundleVersion)

	return nil
}
