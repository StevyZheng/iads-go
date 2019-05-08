package cmd

import (
	"github.com/spf13/cobra"
	"iads/lib/base"
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
		ssh := base.NewSsh("www.roycom.com.cn", "root", "roycom000000")
		_ = ssh.SftpConnect()
		_ = ssh.UploadFile("frp_0.27.0_windows_amd64.zip", "/root/kb.tar.gz")
		//_ = ssh.DownloadFile("/root/kb.tar.gz", "/root/kb.tar.gz")
	},
}
