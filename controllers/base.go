/**
* @Author: kiosk
* @Mail: weijiaxiang007@foxmail.com
* @Date: 2020/3/28
**/
package controllers

import (
	"github.com/astaxie/beego"
	"github.com/tidwall/gjson"
	"imchat-api/global"
	"strings"
)

type BaseController struct {
	beego.Controller
	userName  string
	timestamp string
}

type ReturnResult struct {
	Data   interface{} `json:"data"`
	Result string      `json:"result"`
}

func (this *BaseController) GetParam(key string, canEmpty bool) string {
	ret := this.GetString(key)
	if ret == "" {
		ret = gjson.Get(string(this.Ctx.Input.RequestBody), key).String()
		ret = strings.TrimSpace(ret)
		if ret == "" && canEmpty != true {
			this.Abort("400")
		}
	}
	return ret
}

func (this *BaseController) ReturnError(error string) {
	global.Logger.Errorln(error)

	this.Data["json"] = ReturnResult{error, "error"}
	this.ServeJSON()
	this.Abort("200")
}
