package cmd

import (
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"iads/lib"
	"os"
	"strconv"
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
		arr := lib.Analysis("haha.txt")
		arrLen := arr.Size()
		var data [][]string
		it := arr.Iterator()
		for it.Next() {
			row := make([]string, 2, 2)
			row[0] = strconv.FormatInt(it.Value().(*lib.RowLog).Index, 10)
			row[1] = it.Value().(*lib.RowLog).Data
			data = append(data, row)
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"index", "errMsg"})
		table.SetRowLine(true)
		table.SetFooter([]string{"Total", strconv.Itoa(arrLen)})
		for _, v := range data {
			table.Append(v)
		}
		table.Render()
	},
}
