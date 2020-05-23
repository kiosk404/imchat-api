package main

import (
	"github.com/astaxie/beego"
	"imchat-api/models"
	_ "imchat-api/routers"
)

func init() {
	models.InitMySql()
	//beego.BConfig.WebConfig.Session.SessionOn = true
	//beego.BConfig.WebConfig.Session.SessionProvider="file"
	//beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"

}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
