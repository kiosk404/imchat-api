/**
* @Author: kiosk
* @Mail: weijiaxiang007@foxmail.com
* @Date: 2020/5/16
**/
package controllers

import "imchat-api/services"

type AttachController struct {
	BaseController
}

var uploadService services.UploadService

// @router /upload [post]
func (this *AttachController) Upload() {
	url, err := uploadService.UploadLocal(this.Ctx.ResponseWriter, this.Ctx.Request)
	if err != nil {
		this.ReturnError(err.Error())
	}

	this.Data["json"] = ReturnResult{url, "success"}
	this.ServeJSON()
}
