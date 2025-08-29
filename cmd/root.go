package cmd

import (
	"context"

	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "terraform-mcp-server",
	Short:         "Terraform MCP Server",
	Long:          "Terraform MCP Server - Provides Terraform module and provider documentation via MCP protocol",
	SilenceErrors: true,
}

func init() {}

func Execute(ctx context.Context) error {
	// Remove help for root command
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})

	// Remove shell completion
	rootCmd.CompletionOptions = cobra.CompletionOptions{
		DisableDefaultCmd:   true,
		DisableNoDescFlag:   true,
		DisableDescriptions: true,
		HiddenDefaultCmd:    true,
	}

	return fang.Execute(ctx, rootCmd)
}
