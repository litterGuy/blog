package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type UserController struct {
	beego.Controller
}

// @Title UserInfo
// @Description get user info by id
// @Param   id     query    int  true     "the user`s id"
// @Success 200 {object} models.user
// @Failure 400 Invalid id supplied
// @Failure 404 User not found
//@router /info/:id [get]
func (user *UserController) UserInfo() {
	logs.Info("start handle the request")
	var id, _ = user.GetInt(":id")
	info, rst := (&models.User{}).FindById(id)
	if !rst {
		user.Data["json"] = "has no data"
	} else {
		//此处的"json"为固定死的key值，不以json为key那么记过不会以json格式返回结果
		user.Data["json"] = &info
	}

	ui := models.GetById(id)
	logs.Error(ui)
}

//@router /testError [get]
func (user * UserController) TestError()  {
	user.Abort("500")
}
