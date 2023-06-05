package cmd

import (
	"city-chef/cmd/cook"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

const (
	projectName = "city-chef"
)

func Execute() {
	var rootCmd = &cobra.Command{Use: projectName}
	rootCmd.AddCommand(cook.Command())
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
