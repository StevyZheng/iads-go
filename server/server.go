package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	router2 "iads/server/router"
	"net/http"
	"runtime"
	"time"
)

func ServerStart() {
	runtime.GOMAXPROCS(2)
	gin.SetMode(gin.ReleaseMode)
	router := router2.InitRouter()

	ser := &http.Server{
		Addr:           fmt.Sprintf(":%d", 80),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	_ = ser.ListenAndServe()
}
