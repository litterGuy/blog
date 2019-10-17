package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.Data["json"] = map[string]interface{}{"code":c.Ctx.Output.Status,"errMsg":"request not found"}
	c.ServeJSON()
}

func (c *ErrorController) Error500()  {
	c.Data["json"] = map[string]interface{}{"code":c.Ctx.Output.Status,"errMsg":"server error"}
	c.ServeJSON()
}
