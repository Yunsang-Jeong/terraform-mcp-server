package cmd

import (
	"github.com/Yunsang-Jeong/terraform-mcp-server/pkg/server"
	"github.com/spf13/cobra"
)

var stdioCmd = &cobra.Command{
	Use:   "stdio",
	Short: "Run MCP server over stdio",
	Long: `Run the Terraform MCP server over stdio for direct communication.
This mode is typically used when the server is invoked by an MCP client
that communicates via standard input and output streams.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return server.RunStdio()
	},
}

func init() {
	rootCmd.AddCommand(stdioCmd)
}
