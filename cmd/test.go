package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"iads/lib"
)

func init(){
	rootCmd.AddCommand(testCmd)
}
var testCmd = &cobra.Command{
	Use: "test",
	Short: "test",
	Run: func(cmd *cobra.Command, args []string) {
		tmp := lib.SearchSplitString("hahaha123hehehe4    59 6hshb3io\ntreuhc456j\njwdhfb", ".*h[a-f][0-9].*|wd.*", "[0-9]")
		//tmp := lib.Trim("     he ha     ", " ")
		fmt.Print(tmp)
		fmt.Print(len(tmp))
	},
}