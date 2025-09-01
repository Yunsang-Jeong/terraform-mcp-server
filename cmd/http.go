package cmd

import (
	"github.com/Yunsang-Jeong/terraform-mcp-server/pkg/server"

	"github.com/spf13/cobra"
)

var (
	port uint16
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Run mcp http server",
	RunE: func(cmd *cobra.Command, args []string) error {
		return server.RunHttp(port)
	},
}

func init() {
	httpCmd.Flags().Uint16VarP(&port, "port", "p", 8080, "port to listen on")
	rootCmd.AddCommand(httpCmd)
}
