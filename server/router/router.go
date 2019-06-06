package router

import (
	"github.com/gin-gonic/gin"
	"iads/server/api_v1_0"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	v1_0 := router.Group("/v1.0")

	v1_0.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"app":     "iads",
			"version": "v1.0",
		})
	})
	v1_0.GET("/userlist", api_v1_0.UserList)
	v1_0.POST("/user", api_v1_0.AddUser)
	v1_0.PUT("/user/:id", api_v1_0.UpdateUser)
	v1_0.DELETE("/user/:id", api_v1_0.DeleteUser)

	v1_0.GET("/cpuinfo", api_v1_0.CpuInfo)

	return router
}
