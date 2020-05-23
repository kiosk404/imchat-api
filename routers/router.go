package routers

import (
	"github.com/astaxie/beego"
	"imchat-api/controllers"
	"imchat-api/global"
)

func init() {
	global.Logger.Infoln("Start Init Router")
	beego.Include(&controllers.MainTplController{})
	beego.ErrorController(&controllers.ErrorController{})

	beego.Router("/ws/v1", &controllers.WebSocketController{})
	beego.SetStaticPath("/swagger", "swagger")
	//beego.SetStaticPath("/static","")
	beego.SetStaticPath("/mnt", "mnt")
	ns :=
		beego.NewNamespace("/v1",
			beego.NSNamespace("/attach",
				beego.NSInclude(
					&controllers.AttachController{},
				),
			),
			beego.NSNamespace("/basic",
				beego.NSInclude(
					&controllers.BasicFuncController{},
				),
			),
			beego.NSNamespace("/user",
				beego.NSInclude(
					&controllers.UserController{},
				),
			),
			beego.NSNamespace("/contact",
				beego.NSInclude(
					&controllers.ContactController{},
				),
			),
			beego.NSNamespace("/chat",
				beego.NSInclude(
					&controllers.ChatController{},
				),
			),
		)
	beego.AddNamespace(ns)
}
