/**
* @Author: kiosk
* @Mail: weijiaxiang007@foxmail.com
* @Date: 2020/5/16
**/
package controllers

import (
	"imchat-api/models/template"
	"imchat-api/services"
	"imchat-api/utils"
	"strconv"
)

type ContactController struct {
	BaseController
}

var contactService services.ContactService

type UserListJson struct {
	UserList []services.UserInfoJson
	Total    int
}

type CommunityListJson struct {
	CommunityList []template.Community
	Total         int
}

// @Title Load Friend
// @Description load friend
// @Success 200 {object} services.UserInfoJson
// @Failure 500 backend error
// @router /loadfriend [post]
func (this *ContactController) LoadFriend() {
	var arg services.ContactArg

	if userId, err := strconv.ParseInt(this.Ctx.Request.Form.Get("userid"), 10, 64); err == nil {
		arg.Userid = userId
	} else {
		this.ReturnError(err.Error())
	}

	users, err := contactService.SearchFriend(arg.Userid)
	if err != nil {
		this.ReturnError(err.Error())
	}

	usersJson := UserListJson{UserList: users, Total: len(users)}
	this.Data["json"] = ReturnResult{usersJson, "success"}
	this.ServeJSON()
}

// @Title Add Friend
// @Description add friend
// @Success 200 {object} services.UserInfoJson
// @Failure 500 backend error
// @router /addfriend [post]
func (this *ContactController) AddFriend() {
	var arg services.ContactArg

	userId, err1 := strconv.ParseInt(this.Ctx.Request.Form.Get("userid"), 10, 64)
	dstId, err2 := strconv.ParseInt(this.Ctx.Request.Form.Get("dstid"), 10, 64)
	if err1 != nil || err2 != nil {
		this.ReturnError("参数错误")
	}

	arg.Userid = userId
	arg.Dstid = dstId

	//调用service
	if err := contactService.AddFriend(arg.Userid, arg.Dstid); err != nil {
		this.ReturnError(err.Error())
	}

	this.Data["json"] = ReturnResult{"添加成功", "success"}
	this.ServeJSON()
}

// @Title Load Community
// @Description load community
// @Success 200 {object} services.UserInfoJson
// @Failure 500 backend error
// @router /loadcommunity [post]
func (this *ContactController) LoadCommunity() {
	var arg services.ContactArg

	if userId, err := strconv.ParseInt(this.Ctx.Request.Form.Get("userid"), 10, 64); err == nil {
		arg.Userid = userId
	} else {
		this.ReturnError(err.Error())
	}

	coms, err := contactService.SearchCommunity(arg.Userid)
	if err != nil {
		this.ReturnError(err.Error())
	}
	communitysJson := CommunityListJson{coms, len(coms)}
	this.Data["json"] = ReturnResult{communitysJson, "success"}
	this.ServeJSON()
}

// @Title Create Community
// @Description create community
// @Success 200 {object} controllers.ReturnResult
// @Failure 500 backend error
// @router /createcommunity [post]
func (this *ContactController) CreateCommunity() {
	var community template.Community

	if err := utils.Bind(this.Ctx.Request, &community); err != nil {
		this.ReturnError(err.Error())
	}

	coms, err := contactService.CreateCommunity(community)
	if err != nil {
		this.ReturnError(err.Error())
	}

	this.Data["json"] = ReturnResult{coms, "success"}
	this.ServeJSON()
}



