package user

import "github.com/gin-gonic/gin"

// RegisterRouter 注册路由
func RegisterRouter(r *gin.RouterGroup) {

	// 注册
	r.POST("/register", register)
	// 登录
	r.POST("/login", Auth.LoginHandler)

	auth := r.Group("")
	auth.Use(Auth.MiddlewareFunc())
	{
		// 用户列表
		auth.GET("/list", UserList)
		//添加用户
		auth.POST("/add", AddUser)
		// 删除用户
		auth.DELETE("/:id", DeleteUserByID)
		// 更新用户信息
		auth.PUT("/:id", UpdateUserByID)
	}
}
