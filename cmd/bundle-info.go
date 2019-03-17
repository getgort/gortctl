package cmd

import (
	"fmt"
	"strings"

	"github.com/clockworksoul/cog2/client"
	"github.com/clockworksoul/cog2/data"
	"github.com/spf13/cobra"
)

const (
	bundleInfoUse   = "info"
	bundleInfoShort = "Info a bundle"
	bundleInfoLong  = `
Display bundle information.

If only a bundle name is provided, information on the bundle as a whole is
presented. If that bundle is also currently enabled, details about the
version that is currently live is also displayed.

If a version is also provided, details on that specific version are
presented, regardless of whether it happens to also be enabled.
`
)

// GetBundleInfoCmd is a command
func GetBundleInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   bundleInfoUse,
		Short: bundleInfoShort,
		Long:  bundleInfoLong,
		RunE:  bundleInfoCmd,
		Args:  cobra.RangeArgs(1, 2),
	}

	return cmd
}

func bundleInfoCmd(cmd *cobra.Command, args []string) error {
	switch len(args) {
	case 1:
		return doBundleInfoAll(args[0])
	case 2:
		return doBundleInfoVersion(args[0], args[1])
	}

	return nil
}

func doBundleInfoAll(name string) error {
	cogClient, err := client.Connect(FlagCogProfile)
	if err != nil {
		return err
	}

	bundles, err := cogClient.BundleListVersions(name)
	if err != nil {
		return err
	}

	var enabled *data.Bundle
	var versions = make([]string, 0)

	for _, bundle := range bundles {
		versions = append(versions, bundle.Version)

		if bundle.Enabled {
			enabled = &bundle
		}
	}

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Versions: %s\n", strings.Join(versions, ", "))

	if enabled != nil {
		fmt.Println("Status: Enabled")
		fmt.Printf("Enabled Version: %s\n", enabled.Version)

		commands := make([]string, 0)
		for name := range enabled.Commands {
			commands = append(commands, name)
		}

		fmt.Printf("Commands: %s\n", strings.Join(commands, ", "))
		fmt.Printf("Permissions: %s\n", strings.Join(enabled.Permissions, ", "))
	} else {
		fmt.Println("Status: Disabled")
	}

	return nil
}

func doBundleInfoVersion(name, version string) error {
	cogClient, err := client.Connect(FlagCogProfile)
	if err != nil {
		return err
	}

	bundle, err := cogClient.BundleGet(name, version)
	if err != nil {
		return err
	}

	fmt.Printf("Name: %s\n", bundle.Name)
	fmt.Printf("Version: %s\n", bundle.Version)

	if bundle.Enabled {
		fmt.Println("Status: Enabled")
	} else {
		fmt.Println("Status: Enabled")
	}

	commands := make([]string, 0)
	for name := range bundle.Commands {
		commands = append(commands, name)
	}

	fmt.Printf("Commands: %s\n", strings.Join(commands, ", "))
	fmt.Printf("Permissions: %s\n", strings.Join(bundle.Permissions, ", "))

	return nil
}
