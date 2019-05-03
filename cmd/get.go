package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"iads/lib"
)

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getCpuInfoCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get info",
}

var getCpuInfoCmd = &cobra.Command{
	Use:   "cpuinfo",
	Short: "Print the cpu info",
	Run: func(cmd *cobra.Command, args []string) {
		cpuinfo := new(lib.CpuInfo)
		cpuinfo.GetCpuInfo()
		fmt.Println(cpuinfo.Model)
	},
}
