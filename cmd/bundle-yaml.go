package cmd

import (
	"fmt"

	"github.com/clockworksoul/cog2/client"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

const (
	bundleYamlUse   = "yaml"
	bundleYamlShort = "Retrieve the raw YAML for a bundle."
	bundleYamlLong  = "Retrieve the raw YAML for a bundle."
)

// GetBundleYamlCmd is a command
func GetBundleYamlCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   bundleYamlUse,
		Short: bundleYamlShort,
		Long:  bundleYamlLong,
		RunE:  bundleYamlCmd,
		Args:  cobra.ExactArgs(2),
	}

	return cmd
}

func bundleYamlCmd(cmd *cobra.Command, args []string) error {
	name := args[0]
	version := args[1]

	// TODO Implement that no specifed version returns enabled version.

	cogClient, err := client.Connect(FlagCogProfile)
	if err != nil {
		return err
	}

	bundle, err := cogClient.BundleGet(name, version)
	if err != nil {
		return err
	}

	bytes, err := yaml.Marshal(bundle)
	if err != nil {
		return err
	}

	fmt.Println(string(bytes))

	return nil
}
