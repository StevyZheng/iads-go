package user

import (
	"fmt"
	orm "iads/server/common"
)

var dbConnect *orm.Connection

func init() {
	dbConnect = orm.NewConnection()
}

//用户类
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Phone    string `json:"phone"`
	Role     string `json:"role" form:"role"`
}

func CreateTable() {
	dbConnect.Eloquent.AutoMigrate(User{})
}

var Users []User

type login struct {
	Username string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required"`
}

// Validator .
func (login *login) validator() (*User, string, bool) {
	user := &User{Username: login.Username}
	err := dbConnect.Eloquent.First(&user).Error
	var msg string
	if err != nil {
		msg = "没有该账户！"
		return nil, msg, false
	}

	if user.Password != login.Password {
		msg = "密码错误！"
		fmt.Println("username:", user.Username, "|pwd:", user.Password, "| loginuser:", login.Username, "|loginPwd:", login.Password)
		return nil, msg, false
	}
	msg = "登录成功！"
	return user, msg, true
}

func (user *User) GetOneByUsername(username string) error {
	err = dbConnect.Eloquent.Find(&user).Error
	return err
}

//添加用户
func (user User) Insert() (id int64, err error) {
	ret := dbConnect.Eloquent.Create(&user)
	id = user.ID
	if ret.Error != nil {
		err = ret.Error
		return
	}
	return
}

//用户列表
func (user *User) UserList() (users []User, err error) {
	if err = dbConnect.Eloquent.Find(&users).Error; err != nil {
		return
	}
	return
}

//修改
func (user *User) Update(id int64) (updateUser User, err error) {
	if err = dbConnect.Eloquent.Select([]string{"id", "username"}).First(&updateUser, id).Error; err != nil {
		return
	}
	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = dbConnect.Eloquent.Model(&updateUser).Update(&user).Error; err != nil {
		return
	}
	return
}

//删除数据
func (user *User) Destroy(id int64) (Result User, err error) {
	if err = dbConnect.Eloquent.Select([]string{"id"}).First(&user, id).Error; err != nil {
		return
	}
	if err = dbConnect.Eloquent.Delete(&user).Error; err != nil {
		return
	}
	Result = *user
	return
}
