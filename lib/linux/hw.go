package linux

import (
	"fmt"
	"iads/lib/base"
	"io/ioutil"
	"strings"
)

type CpuInfo struct {
	Model     string
	Count     int
	CoreCount int
	Stepping  string
}

func (e *CpuInfo) GetCpuInfo() {
	//tmpStr, err := ExecShellLinux("cat /proc/cpuinfo")
	tmp, err := ioutil.ReadFile("/proc/cpuinfo")
	tmpStr := strings.Replace(string(tmp), "\n", "", 1)
	if err != nil {
		fmt.Println(err)
	}
	e.Model = base.SearchSplitStringColumnFirst(tmpStr, ".*model name.*", ":", 2)
	e.Stepping = base.SearchSplitStringColumnFirst(tmpStr, ".*stepping.*", ":", 2)
	countTmp1 := base.SearchString(tmpStr, ".*physical id.*")
	countTmp := base.UniqStringList(countTmp1)
	e.Count = len(countTmp)
	coreCountTmp1 := base.SearchString(tmpStr, ".*core id.*")
	coreCountTmp := base.UniqStringList(coreCountTmp1)
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
	tmpStr, err := base.ExecShellLinux("dmidecode")
	if err != nil {
		fmt.Println(err)
	}
	e.Model = base.SearchSplitStringColumnFirst(tmpStr, ".*Product Name.*", ":", 2)
	e.BiosVer = base.SearchSplitStringColumnFirst(tmpStr, ".*Version.*", ":", 2)
	e.BiosDate = base.SearchSplitStringColumnFirst(tmpStr, ".*Release Date.*", ":", 2)
}
