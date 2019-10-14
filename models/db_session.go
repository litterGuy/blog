package models

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)
import "github.com/gocraft/dbr"

var DbSession *dbr.Session

/*
初始化数据连接池
*/
func init() {
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	urls := beego.AppConfig.String("mysqlurls")
	port := beego.AppConfig.String("mysqlport")
	dbs := beego.AppConfig.String("mysqldb")

	conn, err := dbr.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pwd, urls, port, dbs), nil)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()
	if err != nil {
		panic(err)
	}

	//设置数据连接池属性
	conn.SetMaxIdleConns(60)//默认值是2
	conn.SetMaxOpenConns(200)//默认限制是0，表示最大打开数没有限制

	DbSession = conn.NewSession(nil)
}
