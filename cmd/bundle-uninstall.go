package cmd

import (
	"fmt"

	"github.com/getgort/gort/client"
	"github.com/spf13/cobra"
)

const (
	bundleUninstallUse   = "uninstall"
	bundleUninstallShort = "Uninstall a bundle"
	bundleUninstallLong  = "Uninstall a bundle."
)

// GetBundleUninstallCmd is a command
func GetBundleUninstallCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   bundleUninstallUse,
		Short: bundleUninstallShort,
		Long:  bundleUninstallLong,
		RunE:  bundleUninstallCmd,
		Args:  cobra.ExactArgs(2),
	}

	return cmd
}

func bundleUninstallCmd(cmd *cobra.Command, args []string) error {
	bundleName := args[0]
	bundleVersion := args[1]

	c, err := client.Connect(FlagGortProfile)
	if err != nil {
		return err
	}

	err = c.BundleUninstall(bundleName, bundleVersion)
	if err != nil {
		return err
	}

	fmt.Printf("Bundle %q uninstalled.\n", bundleName)

	return nil
}
