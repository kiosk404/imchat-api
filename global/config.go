/**
* @Author: kiosk
* @Mail: weijiaxiang007@foxmail.com
* @Date: 2020/3/28
**/
package global

import (
	"github.com/astaxie/beego"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Config = ReadMFEYaml()

type Environments struct {
	Mysql    MysqlStruct
	LogLevel string
}

type MysqlStruct struct {
	Host     string
	Database string
	Username string
	Password string
}

type MFE struct {
	LocalDev Environments
	Dev      Environments
}

func ReadMFEYaml() Environments {
	data, _ := ioutil.ReadFile("./conf/imchat-api.yml")

	var t = MFE{}
	if err := yaml.Unmarshal(data, &t); err != nil {
		panic("Read MFE Conf Error !")
	}

	env := Environments{}
	runenv := beego.AppConfig.String("cluster")

	switch runenv {
	case "localdev":
		env = t.LocalDev
	case "dev":
		env = t.Dev
	default:
		panic("Read ImChat-API Environments Error !")
	}

	Logger.Info("Read ImChat-API Config:", env)
	return env
}
