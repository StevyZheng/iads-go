package cmd

import (
	"fmt"
	"github.com/jedib0t/go-pretty/table"
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
		d.Run()
		//x := linux.NetInfo{}
		//x.Init()
	},
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run roycom common server test",
	Run: func(cmd *cobra.Command, args []string) {
		tw := table.NewWriter()
		tw.AppendHeader(table.Row{"#", "haha", "hehe", "heihei"})
		tw.AppendRows([]table.Row{
			{1, "1", "2", "hehe3333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333"},
			{2, "1", "2", 3},
		})
		tw.AppendFooter(table.Row{"", "", "foot", 10000})
		tw.SetIndexColumn(1)
		tw.SetAllowedColumnLengths([]int{20, 20, 20, 20})
		fmt.Println(tw.Render())
	},
}
