/**
* @Author: kiosk
* @Mail: weijiaxiang007@foxmail.com
* @Date: 2020/3/28
**/
package controllers

type BasicFuncController struct {
	BaseController
}

// @router /ping [get]
func (this *BasicFuncController) Base() {
	this.Ctx.WriteString("ok")
}
