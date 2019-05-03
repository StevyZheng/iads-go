package cmd

import (
	"github.com/spf13/cobra"
	"iads/lib"
	"runtime"
	"sync"
)

func init(){
	rootCmd.AddCommand(burnCmd)
}

func burnFunc(wg *sync.WaitGroup)  {
	lib.Gaos()
	(*wg).Done()
}

var burnCmd = &cobra.Command{
	Use: "burn",
	Short: "burn cpu: 100%",
	Run: func(cmd *cobra.Command, args []string) {
		runtime.GOMAXPROCS(4)
		var wg sync.WaitGroup
		for i:=0; i<4; i++ {
			wg.Add(1)
			go burnFunc(&wg)
		}
		wg.Wait()
	},
}