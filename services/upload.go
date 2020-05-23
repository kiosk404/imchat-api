/**
* @Author: kiosk
* @Mail: weijiaxiang007@foxmail.com
* @Date: 2020/5/16
**/
package services

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

func init() {
	if err := os.MkdirAll("./mnt", os.ModePerm); err != nil {
		panic("mkdir mnt error: " + err.Error())
	}
}

type UploadService struct {
}

func (up *UploadService) UploadLocal(writer http.ResponseWriter, request *http.Request) (url string, err error) {
	srcfile, head, err := request.FormFile("file")
	if err != nil {
		return url, err
	}

	suffix := ".png"
	//如果前端文件名称包含后缀 xx.xx.png
	ofilename := head.Filename
	tmp := strings.Split(ofilename, ".")
	if len(tmp) > 1 {
		suffix = "." + tmp[len(tmp)-1]
	}
	//如果前端指定filetype
	filetype := request.FormValue("filetype")
	if len(filetype) > 0 {
		suffix = filetype
	}
	//time.Now().Unix()
	filename := fmt.Sprintf("%d%04d%s", time.Now().Unix(), rand.Int31(), suffix)
	dstfile, err := os.Create("./mnt/" + filename)
	if err != nil {
		return url, err
	}
	// 将源文件内容copy到新文件
	_, err = io.Copy(dstfile, srcfile)
	if err != nil {
		return url, err
	}
	// 将新文件路径转换成url地址
	url = "/mnt/" + filename
	return url, nil
}
