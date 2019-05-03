package lib

import (
	"bytes"
	"os/exec"
)

func ExecShellLinuxE(cmd string) (string, error)  {
	ret := exec.Command("/bin/bash", "-c", cmd)
	var out bytes.Buffer
	ret.Stdout = &out
	err := ret.Run()
	return out.String(), err
}

type CpuInfo struct {
	Model string
	Count int
	CoreCount int
	Stepping string
}

func GetCpuInfo(e CpuInfo){
	tmpStr, err := ExecShellLinux("cat /proc/cpuinfo")
	if err != nil{
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