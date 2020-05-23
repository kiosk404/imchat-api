/**
* @Author: kiosk
* @Mail: weijiaxiang007@foxmail.com
* @Date: 2020/3/28
**/
package global

import (
	log "github.com/sirupsen/logrus"
)

var Logger = log.New()

func init() {
	switch Config.LogLevel {
	case "debug":
		Logger.SetLevel(log.DebugLevel)
	case "info":
		Logger.SetLevel(log.InfoLevel)
	case "warn":
		Logger.SetLevel(log.WarnLevel)
	case "fatal":
		Logger.SetLevel(log.FatalLevel)
	case "panic":
		Logger.SetLevel(log.PanicLevel)
	default:
		Logger.SetLevel(log.ErrorLevel)
	}
	Logger.SetReportCaller(false)
	Logger.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
}
