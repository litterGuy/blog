package filter

import (
	"github.com/astaxie/beego/context"
)

/*
对返回json结果进行统一处理
*/
func NsJsonFilter(ctx *context.Context) {
	rst := ctx.Input.GetData("json")

	//json.MarshalIndent 不支持map[interface {}]interface {}类型，故在此处修改key为string
	packageRst := make(map[string]interface{})
	if _, ok := rst.(string); ok {
		packageRst["code"] = 1
		packageRst["errMsg"] = rst
	} else {
		packageRst["code"] = 0
		packageRst["data"] = rst
	}

	//将封装后的数据重新写回body中
	ctx.Output.Reset(ctx)
	ctx.Output.JSON(packageRst, true, false)
}
