package main

import (
	"blog/filter"
	_ "blog/routers"

	"github.com/astaxie/beego"
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
	beego.InsertFilter("/*", beego.FinishRouter, filter.NsJsonFilter,false)

	beego.Run()
}
