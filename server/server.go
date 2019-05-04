package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ServerStart() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"app": "iads",
			"ver": "1.0.0",
		})
	})

	ser := &http.Server{
		Addr:           fmt.Sprintf(":%d", 8080),
		Handler:        router,
		ReadTimeout:    60,
		WriteTimeout:   60,
		MaxHeaderBytes: 1 << 20,
	}
	_ = ser.ListenAndServe()
}
