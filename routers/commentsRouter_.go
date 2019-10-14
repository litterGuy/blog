package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["blog/controllers:UserController"] = append(beego.GlobalControllerRouter["blog/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserInfo",
            Router: `/info/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
