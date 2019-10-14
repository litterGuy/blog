package models

import (
	"fmt"
	"github.com/astaxie/beego"
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

func init() {
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	urls := beego.AppConfig.String("mysqlurls")
	port := beego.AppConfig.String("mysqlport")
	dbs := beego.AppConfig.String("mysqldb")

	// set default database
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pwd, urls, port, dbs), 30)

	// register model
	orm.RegisterModel(new(User))

	// create table
	orm.RunSyncdb("default", false, true)
}
