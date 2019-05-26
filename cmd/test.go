package cmd

import (
	"flag"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
	"time"

	//"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
	"iads/lib/linux"
	_ "iads/server/api_v1_0"
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
		var kubeconfig *string
		if home := homeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
		flag.Parse()
		// uses the current context in kubeconfig
		config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			panic(err.Error())
		}
		// creates the clientset
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			panic(err.Error())
		}
		for {
			pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
			time.Sleep(10 * time.Second)
		}
	},
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
