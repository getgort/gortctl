package cmd

import (
	"github.com/spf13/cobra"
)

const (
	bundleUse   = "bundle"
	bundleShort = "Perform operations on bundles"
	bundleLong  = "Allows you to perform bundle administration."
)

// $ cogctl bundle --help
// Usage: cogctl bundle [OPTIONS] COMMAND [ARGS]...

//   Manage command bundles and their config.

//   If no subcommand is given, lists all bundles installed, and their
//   currently enabled version, if any.

// Options:
//   -e, --enabled   List only enabled bundles
//   -d, --disabled  List only disabled bundles
//   -v, --verbose   Display additional bundle details
//   --help          Show this message and exit.

// Commands:
//   config     Manage dynamic configuration layers.
//   disable    Disable a bundle by name.
//   enable     Enable the specified version of the bundle.
//   info       Display bundle information.
//   install    Install a bundle.
//   uninstall  Uninstall bundles.
//   versions   List installed bundle versions.

// GetBundleCmd bundle
func GetBundleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   bundleUse,
		Short: bundleShort,
		Long:  bundleLong,
	}

	cmd.AddCommand(GetBundleInstallCmd())
	cmd.AddCommand(GetBundleListCmd())

	return cmd
}
