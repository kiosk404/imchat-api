/**
* @Author: kiosk
* @Mail: weijiaxiang007@foxmail.com
* @Date: 2020/5/16
**/
package controllers

import (
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"imchat-api/global"
	"imchat-api/services"
	"net/http"
	"strconv"
)

type ChatController struct {
	BaseController
}

type WebSocketController struct {
	BaseController
}

var chatService services.ChatService

func (this *WebSocketController) Get() {
	query := this.Ctx.Request.URL.Query()
	id := query.Get("id")
	token := query.Get("token")

	userId, _ := strconv.ParseInt(id, 10, 64)
	isvalida := userService.CheckToken(userId, token)

	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return isvalida
		},
	}).Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)

	if err != nil {
		global.Logger.Error("Error : " + err.Error())
		return
	}
	// 获得conn
	node := &services.Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 500),
		GroupSets: set.New(set.ThreadSafe),
	}
	// 获取用户全部群Id
	comIds := contactService.SearchCommunityIds(userId)
	for _, v := range comIds {
		node.GroupSets.Add(v)
	}
	// userid 和 node 形成绑定关系
	services.RWLocker.Lock()
	services.ClientMap[userId] = node
	services.RWLocker.Unlock()

	// 发送逻辑
	go chatService.SendProc(node)
	// 接收逻辑
	go chatService.RecvProc(node)

	global.Logger.Infof("<- %d\n", userId)

	chatService.SendMsg(userId, []byte("hello,world!"))
}
