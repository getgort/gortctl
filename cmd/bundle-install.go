package cmd

import (
	"fmt"

	"github.com/clockworksoul/cog2/bundle"
	"github.com/clockworksoul/cog2/client"
	"github.com/spf13/cobra"
)

const (
	bundleInstallUse   = "install"
	bundleInstallShort = "Install a bundle"
	bundleInstallLong  = "Install a bundle."
)

// GetBundleInstallCmd is a command
func GetBundleInstallCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   bundleInstallUse,
		Short: bundleInstallShort,
		Long:  bundleInstallLong,
		RunE:  bundleInstallCmd,
		Args:  cobra.ExactArgs(1),
	}

	return cmd
}

func bundleInstallCmd(cmd *cobra.Command, args []string) error {
	bundlefile := args[0]

	c, err := client.Connect(FlagCogProfile)
	if err != nil {
		return err
	}

	bundle, err := bundle.LoadBundle(bundlefile)
	if err != nil {
		return err
	}

	err = c.BundleInstall(bundle)
	if err != nil {
		return err
	}

	fmt.Printf("Bundle %q installed.\n", bundle.Name)

	return nil
}
