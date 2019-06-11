package user

import (
	"github.com/gin-gonic/gin"
	"iads/server/common"
	"iads/server/common/pkg/e"
	"net/http"
	"strconv"
)

// @Summary 用户注册
// @Produce  json
// @Param name body string true "Name"
// @Param state body int false "State"
// @Param created_by body int false "CreatedBy"
// @Success 200
// @Failure 500
// @Router /api/v1/tags [post]
func register(c *gin.Context) {
	u := &User{}
	if err := c.ShouldBindJSON(u); err != nil {
		common.RES(c, e.INVALID_PARAMS, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err = u.GetOneByUsername(u.Username); err != nil {
		_, err := u.Insert()
		if err != nil {
			common.RES(c, e.ERROR, gin.H{
				"message": err.Error(),
			})
		} else {
			common.RES(c, e.SUCCESS, gin.H{})
		}
	} else {
		common.RES(c, e.ERROR, gin.H{
			"message": "用户名已存在！",
		})
	}
}

//列表数据
func UserList(c *gin.Context) {
	var users User
	result, err := users.UserList()
	if err != nil {
		common.RES(c, e.ERROR, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		common.RES(c, e.ERROR, gin.H{
			"data": result,
		})
	}
}

// @Summary 添加用户
// @Description add user by username and password
// @Accept  json
// @Produce  json
// @Param  username query string true "Username"
// @Param  password query string true "Password"
// @Success 200 {string} string	"ok"
// @Router /v1.0/useradd [post]
func AddUser(c *gin.Context) {
	var user User
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
func UpdateUserByID(c *gin.Context) {
	var user User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	user.Password = c.Request.FormValue("password")
	_, err = user.Update(id)
	if err != nil {
		common.RES(c, e.ERROR, gin.H{
			"message": err.Error(),
		})
	} else {
		common.RES(c, e.SUCCESS, gin.H{})
	}
}

//删除数据
func DeleteUserByID(c *gin.Context) {
	var user User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if id == 0 {
		common.RES(c, e.INVALID_PARAMS, gin.H{
			"message": "id必须大于0",
		})
	}
	_, err = user.Destroy(id)
	if err != nil {
		common.RES(c, e.ERROR, gin.H{
			"message": err.Error(),
		})
	} else {
		common.RES(c, e.SUCCESS, gin.H{})
	}
}
