package lib

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"log"
	"os"
)

type LogErrMsgStruct struct {
	id       int
	ErrMsg   string
	ErrSolve string
}

var (
	fp     *os.File
	err    error
	logger *log.Logger

	logFiles arraylist.List //日志文件路径数组
	errMsgs  arraylist.List //错误信息数组
)

func InitEnv() {
	//添加日志文件路径
	logFiles.Add("/var/log/messages", "/var/log/mcelog", "/var/log/kerl", "/var/log/syslog")
	errMsgs.Add("error", "failed")

	fp, err = os.Create("testlog.txt")
	if err != nil {
		log.Fatal("create logfile failed.")
	}
	logger = log.New(fp, "", log.LstdFlags|log.Llongfile)
}

func Start() {
	logger.Println("Start testing......")
}

func Stop() {
	logger.Println("test stopped.")
	_ = fp.Sync()
	fp.Close()
}
