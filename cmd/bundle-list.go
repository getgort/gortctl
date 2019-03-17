package cmd

import (
	"fmt"

	"github.com/clockworksoul/cog2/client"
	"github.com/spf13/cobra"
)

const (
	bundleListUse   = "list"
	bundleListShort = "List all existing bundles"
	bundleListLong  = "List all existing bundles."
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

// GetBundleListCmd is a command
func GetBundleListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   bundleListUse,
		Short: bundleListShort,
		Long:  bundleListLong,
		RunE:  bundleListCmd,
	}

	return cmd
}

func bundleListCmd(cmd *cobra.Command, args []string) error {
	const format = "%-12s%-10s%s\n"

	cogClient, err := client.Connect(FlagCogProfile)
	if err != nil {
		return err
	}

	bundles, err := cogClient.BundleList()
	if err != nil {
		return err
	}

	fmt.Printf(format, "BUNDLE", "VERSION", "STATUS")

	for _, b := range bundles {
		status := "Disabled"

		if b.Enabled {
			status = "Enabled"
		}

		fmt.Printf(format, b.Name, b.Version, status)
	}

	return nil
}
