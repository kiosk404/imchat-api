/**
* @Author: kiosk
* @Mail: weijiaxiang007@foxmail.com
* @Date: 2020/5/16
**/
package controllers

type MainTplController struct {
	BaseController
}

// @router / [get]
func (this *MainTplController) Main() {
	this.TplName = "user/index.html"
}

// @router /register [get]
func (this *MainTplController) Register() {
	this.TplName = "user/register.html"
}

// @router /chat [get]
func (this *MainTplController) ChatIndex() {
	this.TplName = "chat/index.html"
}

// @router /createcom [get]
func (this *MainTplController) CreateCommunity() {
	this.TplName = "chat/createcom.html"
}
