package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"iads/server/common/middlewares"
	v1 "iads/server/routers/api/v1"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	rbac := middlewares.RbacInitReturnEnforcer()
	router.Use(sessions.Sessions("default", store))
	router.Use(middlewares.NewAuthorizer(rbac))
	api := router.Group("/api")
	{
		v1.RegisterRouter(api)
	}
	return router
}
