package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"iads/lib"
)

func init() {
	rootCmd.AddCommand(logCmd)
	logCmd.AddCommand(errCmd)
}

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "log functions",
}

var errCmd = &cobra.Command{
	Use:   "err",
	Short: "print err log",
	Run: func(cmd *cobra.Command, args []string) {
		lib.InitEnv()

		arr := lib.Analysis("/var/log/messages")
		fmt.Println("/var/log/messages")
		if 0 == arr.Size() {
			fmt.Println("No errors.")
		} else {
			it := arr.Iterator()
			for it.Next() {
				tRow := it.Value().(*lib.RowLog)
				tStr := fmt.Sprintf("%d %s", tRow.Index, tRow.Data)
				fmt.Println(tStr)
			}
		}

		arr.Clear()
		arr = lib.Analysis("/var/log/mcelog")
		fmt.Println("/var/log/mcelog")
		if 0 == arr.Size() {
			fmt.Println("No errors.")
		} else {
			it := arr.Iterator()
			for it.Next() {
				tRow := it.Value().(*lib.RowLog)
				tStr := fmt.Sprintf("%d %s", tRow.Index, tRow.Data)
				fmt.Println(tStr)
			}
		}

	},
}
