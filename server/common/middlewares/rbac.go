package middlewares

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
	"github.com/gin-gonic/gin"
)

func RbacInitReturnEnforcer() *casbin.Enforcer {
	a := gormadapter.NewAdapter("mysql", "root:000000@tcp(127.0.0.1:3306)/iads?charset=utf8&parseTime=True&loc=Local&timeout=10ms", true)
	e := casbin.NewEnforcer("rbac.conf", a)
	//从DB加载策略
	_ = e.LoadPolicy()
	return e
}

//拦截器
func RbacHandler(e *casbin.Enforcer) gin.HandlerFunc {

	return func(c *gin.Context) {

		//获取请求的URI
		obj := c.Request.URL.RequestURI()
		//获取请求方法
		act := c.Request.Method
		//获取用户的角色
		sub := "admin"

		//判断策略中是否存在
		if e.Enforce(sub, obj, act) {
			fmt.Println("通过权限")
			c.Next()
		} else {
			fmt.Println("权限没有通过")
			c.Abort()
		}
	}
}
