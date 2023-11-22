package cmd

import (
	"github.com/aethiopicuschan/penguin/core"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:               "penguin",
	Long:              "Penguin is a Interactive boilerplate.",
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	Version:           "0.0.1",
	SilenceUsage:      true,
	RunE:              core.Main,
}

func Execute() error {
	return rootCmd.Execute()
}
