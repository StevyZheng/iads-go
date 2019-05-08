package cmd

import (
	"github.com/spf13/cobra"
	"iads/lib/linux"
)

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.AddCommand(commonCmd)
	testCmd.AddCommand(runCmd)
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "run roycom server test",
}

var commonCmd = &cobra.Command{
	Use:   "common",
	Short: "test",
	Run: func(cmd *cobra.Command, args []string) {
		/*ssh := base.NewSsh("www.roycom.com.cn", "root", "roycom000000")
		_ = ssh.SftpConnect()
		_ = ssh.UploadFile("frp_0.27.0_windows_amd64.zip", "/root/kb.tar.gz")
		_ = ssh.DownloadFile("/root/kb.tar.gz", "/root/kb.tar.gz")*/
		d := linux.DmiInfo{}
		d.Run2()
	},
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run roycom common server test",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
