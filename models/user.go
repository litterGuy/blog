package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id    int
	Phone string
	Email string
	Name  string
	Age   int8
	Sex   bool
}

//根据id获取信息
func (user *User) FindById(id int) (*User, bool) {
	count, err := DbSession.Select("*").From("user").Where("id=?", id).Load(user)
	if count <= 0 || err != nil {
		logs.Warn("user find by id %s error and error is %s", id, err)
		return nil, false
	}
	return user, true
}

func GetById(id int) (user *User) {
	u := User{Id: id}
	err := orm.NewOrm().Read(&u)
	if err != nil {
		logs.Error(err)
		return nil
	}
	return &u
}
