package linux

import (
	"encoding/json"
	"fmt"
	"github.com/dselans/dmidecode"
	"github.com/shirou/gopsutil/cpu"
	"iads/lib/base"
	"io/ioutil"
	"strings"
)

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

type CpuHwInfo struct {
	Model     string
	Count     int
	CoreCount int
	Stepping  string
}

func (e *CpuHwInfo) GetCpuHwInfo() {
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

type DmiInfo struct {
}

func (e *DmiInfo) Run() {
	dmi := dmidecode.New()
	if err := dmi.Run(); err != nil {
		fmt.Printf("Unable to get dmidecode information. Error: %v\n", err)
	}
	for handle, record := range dmi.Data {
		fmt.Println("Checking record:", handle)
		for k, v := range record {
			fmt.Printf("Key: %v Val: %v\n", k, v)
		}
	}
}

func (e *DmiInfo) Run2() {
	x, _ := cpu.Info()
	j, _ := json.MarshalIndent(x, "", "    ")
	s := string(j)
	fmt.Println(s)
}
