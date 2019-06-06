package models

import (
	orm "iads/server/database"
)

//用户类
type User struct {
	ID         int64  `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Gender     string `json:"gender"`
	Phone      string `json:"phone"`
	Permission string `json:"permission"`
}

func CreateTable() {
	orm.Eloquent.AutoMigrate(User{})
}

var Users []User

//添加用户
func (user User) Insert() (id int64, err error) {
	ret := orm.Eloquent.Create(&user)
	id = user.ID
	if ret.Error != nil {
		err = ret.Error
		return
	}
	return
}

//用户列表
func (user *User) UserList() (users []User, err error) {
	if err = orm.Eloquent.Find(&users).Error; err != nil {
		return
	}
	return
}

//修改
func (user *User) Update(id int64) (updateUser User, err error) {
	if err = orm.Eloquent.Select([]string{"id", "username"}).First(&updateUser, id).Error; err != nil {
		return
	}
	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Model(&updateUser).Update(&user).Error; err != nil {
		return
	}
	return
}

//删除数据
func (user *User) Destroy(id int64) (Result User, err error) {
	if err = orm.Eloquent.Select([]string{"id"}).First(&user, id).Error; err != nil {
		return
	}
	if err = orm.Eloquent.Delete(&user).Error; err != nil {
		return
	}
	Result = *user
	return
}
