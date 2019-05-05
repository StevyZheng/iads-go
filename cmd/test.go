package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.AddCommand(commonCmd)
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "run roycom server test",
}

var commonCmd = &cobra.Command{
	Use:   "common",
	Short: "run roycom common server test",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
