package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func Execute() {
	// Remove help for root command
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})

	// Remove shell completion
	rootCmd.CompletionOptions = cobra.CompletionOptions{
		DisableDefaultCmd:   true,
		DisableNoDescFlag:   true,
		DisableDescriptions: true,
		HiddenDefaultCmd:    true,
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
