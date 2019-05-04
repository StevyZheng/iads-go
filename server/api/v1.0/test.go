package v1_0

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/json"
	"iads/lib"
	"net/http"
)

func Test(c *gin.Context) {
	cpu := lib.CpuInfo{}
	cpu.GetCpuInfo()
	jsons, errs := json.Marshal(cpu)
	if errs != nil {
	}
	c.String(http.StatusOK, string(jsons))
}
