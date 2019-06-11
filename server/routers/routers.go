package routers

import (
	"github.com/gin-gonic/gin"
	v1 "iads/server/routers/api/v1"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	//router.Use(middlewares.NewAuthorizer(e))
	api := router.Group("/api")
	{
		v1.RegisterRouter(api)
	}
	return router
}
