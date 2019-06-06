package api_v1_0

import (
	"github.com/gin-gonic/gin"
	"iads/server/models"
	"net/http"
	"strconv"
)

type LoginResult struct {
	Token string `json:"token"`
	models.User
}

// 登录
func Login(c *gin.Context) {
	var loginReq module.LoginReq
	if c.BindJSON(&loginReq) == nil {
		isPass, user, err := module.LoginCheck(loginReq)
		if isPass {
			generateToken(c, user)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "验证失败" + err.Error(),
			})
			return
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "json 解析失败",
		})
		return
	}
}

//列表数据
func UserList(c *gin.Context) {
	var user models.User
	user.Username = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")
	result, err := user.UserList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "抱歉，未找到相关信息",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
	})
}

//添加数据
func AddUser(c *gin.Context) {
	var user models.User
	user.Username = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")
	id, err := user.Insert()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "添加失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "添加成功",
		"data":    id,
	})
}

//修改数据
func UpdateUser(c *gin.Context) {
	var user models.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	user.Password = c.Request.FormValue("password")
	result, err := user.Update(id)
	if err != nil || result.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "修改失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "修改成功",
	})
}

//删除数据
func DeleteUser(c *gin.Context) {
	var user models.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	result, err := user.Destroy(id)
	if err != nil || result.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "删除成功",
	})
}
