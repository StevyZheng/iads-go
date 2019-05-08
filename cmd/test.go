package cmd

import (
	"github.com/spf13/cobra"
	"iads/lib"
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
		ssh := lib.NewSsh("192.168.1.111", "root", "000000")
		_ = ssh.SftpConnect()
		_ = ssh.UploadFile("kb.tar.gz", "/root/kb.tar.gz")
	},
}
