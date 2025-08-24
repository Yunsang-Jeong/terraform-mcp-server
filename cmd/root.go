package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "terraform-mcp-server",
	Short:         "terraform-mcp-server",
	Long:          "terraform-mcp-server",
	SilenceErrors: true,
}

func init() {}
