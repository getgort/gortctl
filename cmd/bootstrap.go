package cmd

import (
	"github.com/spf13/cobra"
)

const (
	bootstrapUse   = "bootstrap"
	bootstrapShort = "Bootstrap a new server."
	bootstrapLong  = `Bootstrap can be used on a brand new server to create the administrative user (admin).
	
It will only work on an instance that doesn't yet have any users defined.`
)

// GetBootstrapCmd bootstrap
func GetBootstrapCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   bootstrapUse,
		Short: bootstrapShort,
		Long:  bootstrapLong,
	}

	return cmd
}
