/**
* @Author: kiosk
* @Mail: weijiaxiang007@foxmail.com
* @Date: 2020/3/28
**/
package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"imchat-api/global"
	"imchat-api/models/template"
	"time"
)

func InitMySql() {
	if global.Config.LogLevel == "debug" {
		orm.Debug = true
	} else {
		orm.Debug = false
	}

	global.Logger.Debug(global.Config.Mysql)

	connInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", global.Config.Mysql.Username, global.Config.Mysql.Password, global.Config.Mysql.Host, global.Config.Mysql.Database)
	conn_num := 0

	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		global.Logger.Errorln("Registe MySQL Driver Error: ", err.Error())
		panic(err)
	}
	for {
		err := orm.RegisterDataBase("default", "mysql", connInfo, 30)
		if err != nil {
			conn_num += 1
			time.Sleep(1e9)
			if conn_num > 10 {
				global.Logger.Errorln("Connect MySql Error... connect refuse:", err.Error())
				break
			}
		} else {
			global.Logger.Infoln("Connect MySQL Successed ...")
			break
		}
	}
	maxidelconn, err := beego.AppConfig.Int("MaxIdleConns")
	if err != nil {
		global.Logger.Errorln("Can't read MaxIdleConns from app.conf ...")
		panic(err)
	}
	maxopenconn, err := beego.AppConfig.Int("MaxOpenConns")
	if err != nil {
		global.Logger.Errorln("Can't read MaxOpenConns from app.conf ...")
		panic(err)
	}
	rowslimit, err := beego.AppConfig.Int("DefaultRowsLimit")
	if err != nil {
		global.Logger.Errorln("Can't read DefaultRowsLimit from app.conf ...")
		panic(err)
	}

	orm.DefaultRowsLimit = rowslimit
	orm.SetMaxIdleConns("default", maxidelconn)
	orm.SetMaxOpenConns("default", maxopenconn)

	orm.RegisterModel(
		new(template.Community),
		new(template.User),
		new(template.Contact),
	)
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		global.Logger.Errorln("RunSyncdb Error:", err.Error())
	}
}

func GenerateDBOrm() orm.Ormer {
	return orm.NewOrm()
}
