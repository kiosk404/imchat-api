/**
* @Author: kiosk
* @Mail: weijiaxiang007@foxmail.com
* @Date: 2020/3/28
**/
package controllers

import (
	"imchat-api/models/template"
	"imchat-api/services"
	"imchat-api/utils"
)

type UserController struct {
	BaseController
}

var userService services.UserService

// @Title Get UserInfo
// @Description get user info
// @Success 200 {object} services.UserInfoJson
// @Failure 500 backend error
// @router /info [get]
func (this *UserController) UserInfo() {
	userinfoJson := services.UserInfoJson{}

	this.Data["json"] = ReturnResult{userinfoJson, "success"}
	this.ServeJSON()
}

// @Title UserLogin
// @Description Determine whether the user is logged in
// @Success 200 {object} services.UserInfoJson
// @Failure 500 backend error
// @router /login [post]
func (this *UserController) UserLogin() {
	mobile := this.GetParam("mobile", false)
	passwd := this.GetParam("passwd", false)

	userLoginJson := services.UserInfoJson{}

	user, err := userService.Login(mobile, passwd)
	if err != nil {
		this.ReturnError(err.Error())
	}

	userService.LoadUserInfo(&userLoginJson, user)
	this.Data["json"] = ReturnResult{userLoginJson, "success"}
	this.ServeJSON()
}

// @Title Get UserRegister
// @Description user register
// @Success 200 {object} services.UserInfoJson
// @Failure 500 backend error
// @router /register [post]
func (this *UserController) UserRegister() {
	userRegisterJson := services.UserInfoJson{}
	user := template.User{}

	if err := utils.Bind(this.Ctx.Request, &user); err != nil {
		this.ReturnError(err.Error())
	}

	if err := userService.Register(&user); err != nil {
		this.ReturnError(err.Error())
	}

	userService.LoadUserInfo(&userRegisterJson, user)
	this.Data["json"] = ReturnResult{userRegisterJson, "success"}
	this.ServeJSON()
}

// @Title Get UserFind
// @Description get user find result
// @Success 200 {object} services.UserInfoJson
// @Failure 500 backend error
// @router /find [get]
func (this *UserController) UserFind() {
	userFindJson := services.UserInfoJson{}

	this.Data["json"] = ReturnResult{userFindJson, "success"}
	this.ServeJSON()
}
