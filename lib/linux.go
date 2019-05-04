package lib

import (
	"bytes"
	"fmt"
	"os/exec"
)

type LogErrMsgStruct struct {
	id       int
	ErrMsg   string
	ErrSolve string
}

func ExecShellLinuxE(cmd string) (string, error) {
	ret := exec.Command("/bin/bash", "-c", cmd)
	var out bytes.Buffer
	ret.Stdout = &out
	err := ret.Run()
	return out.String(), err
}

type CpuInfo struct {
	Model     string
	Count     int
	CoreCount int
	Stepping  string
}

func (e *CpuInfo) GetCpuInfo() {
	tmpStr, err := ExecShellLinux("cat /proc/cpuinfo")
	if err != nil {
		fmt.Println(err)
	}
	e.Model = SearchSplitStringColumnFirst(tmpStr, ".*model name.*", ":", 2)
	e.Stepping = SearchSplitStringColumnFirst(tmpStr, ".*stepping.*", ":", 2)
	countTmp1 := SearchString(tmpStr, ".*physical id.*")
	countTmp := UniqStringList(countTmp1)
	e.Count = len(countTmp)
	coreCountTmp1 := SearchString(tmpStr, ".*core id.*")
	coreCountTmp := UniqStringList(coreCountTmp1)
	e.CoreCount = len(coreCountTmp)
}

type MotherboradInfo struct {
	Model    string
	BiosVer  string
	BiosDate string
	BmcVer   string
	BmcDate  string
}

func (e *MotherboradInfo) GetMbInfo() {
	tmpStr, err := ExecShellLinux("dmidecode")
	if err != nil {
		fmt.Println(err)
	}
	e.Model = SearchSplitStringColumnFirst(tmpStr, ".*Product Name.*", ":", 2)
	e.BiosVer = SearchSplitStringColumnFirst(tmpStr, ".*Version.*", ":", 2)
	e.BiosDate = SearchSplitStringColumnFirst(tmpStr, ".*Release Date.*", ":", 2)
}
