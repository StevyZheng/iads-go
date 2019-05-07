package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"iads/server/api_v1_0"
	"net/http"
	"runtime"
	"time"
)

func ServerStart() {
	runtime.GOMAXPROCS(2)
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	v1_0 := router.Group("/v1.0")
	v1_0.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"app":     "iads",
			"version": "v1.0",
		})
	})
	v1_0.GET("/cpuinfo", api_v1_0.CpuInfo)

	ser := &http.Server{
		Addr:           fmt.Sprintf(":%d", 80),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	_ = ser.ListenAndServe()
}
