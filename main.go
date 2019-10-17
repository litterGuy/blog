package main

import (
	"blog/filter"
	"blog/health"
	"blog/models"
	_ "blog/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	/*
		设置 returnOnOutput 的值(默认 true), 如果在进行到此过滤之前已经有输出，
		是否不再继续执行此过滤器,默认设置为如果前面已有输出(参数为true)，则不再执行此过滤器
	*/
	beego.InsertFilter("/*", beego.FinishRouter, filter.NsJsonFilter, false)

	beego.Run()
}

func init() {
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	urls := beego.AppConfig.String("mysqlurls")
	port := beego.AppConfig.String("mysqlport")
	dbs := beego.AppConfig.String("mysqldb")
	// set default database
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pwd, urls, port, dbs), 30, 30)
	// register model
	orm.RegisterModel(new(models.User))
	// create table
	orm.RunSyncdb("default", false, true)
	orm.Debug = true

	toolbox.AddHealthCheck("database", &health.DatabaseCheck{})
}
