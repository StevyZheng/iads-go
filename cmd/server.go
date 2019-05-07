package cmd

import (
	"github.com/spf13/cobra"
	"iads/server"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "restful api_v1_0 server",
	Run: func(cmd *cobra.Command, args []string) {
		server.ServerStart()
	},
}
