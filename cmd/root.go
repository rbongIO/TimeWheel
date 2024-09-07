package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(ant)
	rootCmd.AddCommand(redis_cmd)
	rootCmd.AddCommand(chanRR)
}
