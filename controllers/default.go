package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type DispatchController struct {
	beego.Controller
}

func (this *DispatchController) Get() {
}

type ErrorController struct {
	beego.Controller
}

func (this *ErrorController) Error404() {
	this.TplName = "404.html"
}

func (this *ErrorController) Error500() {
	this.TplName = "500.html"
}
