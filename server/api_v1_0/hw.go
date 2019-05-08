package api_v1_0

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/json"
	"iads/lib/linux"
	"net/http"
)

func CpuInfo(c *gin.Context) {
	cpu := linux.CpuInfo{}
	cpu.GetCpuInfo()
	jsons, errs := json.MarshalIndent(cpu, "", "  ")
	if errs != nil {
	}
	c.String(http.StatusOK, string(jsons))
}
