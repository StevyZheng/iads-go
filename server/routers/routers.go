package routers

import (
	"github.com/gin-gonic/gin"
	"iads/server/common/middlewares"
	v1 "iads/server/routers/api/v1"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	//router.Use(middlewares.NewAuthorizer(e))
	rbac := middlewares.RbacInitReturnEnforcer()
	router.Use(middlewares.RbacHandler(rbac))
	api := router.Group("/api")
	{
		v1.RegisterRouter(api)
	}
	return router
}
