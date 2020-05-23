/**
* @Author: kiosk
* @Mail: weijiaxiang007@foxmail.com
* @Date: 2020/5/23
**/
package services

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"imchat-api/global"
	"net"
	"sync"
)

type ChatService struct {
}

const (
	CMD_SINGLE_MSG = 10
	CMD_ROOM_MSG   = 11
	CMD_HEART      = 0
)

type Message struct {
	Id      int64  `json:"id,omitempty" form:"id"`           //消息ID
	Userid  int64  `json:"userid,omitempty" form:"userid"`   //谁发的
	Cmd     int    `json:"cmd,omitempty" form:"cmd"`         //群聊还是私聊
	Dstid   int64  `json:"dstid,omitempty" form:"dstid"`     //对端用户ID/群ID
	Media   int    `json:"media,omitempty" form:"media"`     //消息按照什么样式展示
	Content string `json:"content,omitempty" form:"content"` //消息的内容
	Pic     string `json:"pic,omitempty" form:"pic"`         //预览图片
	Url     string `json:"url,omitempty" form:"url"`         //服务的URL
	Memo    string `json:"memo,omitempty" form:"memo"`       //简单描述
	Amount  int    `json:"amount,omitempty" form:"amount"`   //其他和数字相关的
}

/**
消息发送结构体
1、MEDIA_TYPE_TEXT
	{id:1,userid:2,dstid:3,cmd:10,media:1,content:"hello"}
2、MEDIA_TYPE_News
	{id:1,userid:2,dstid:3,cmd:10,media:2,content:"标题",pic:"http://www.baidu.com/a/log,jpg",url:"http://www.a,com/dsturl","memo":"这是描述"}
3、MEDIA_TYPE_VOICE，amount单位秒
	{id:1,userid:2,dstid:3,cmd:10,media:3,url:"http://www.a,com/dsturl.mp3",anount:40}
4、MEDIA_TYPE_IMG
	{id:1,userid:2,dstid:3,cmd:10,media:4,url:"http://www.baidu.com/a/log,jpg"}
5、MEDIA_TYPE_REDPACKAGR //红包amount 单位分
	{id:1,userid:2,dstid:3,cmd:10,media:5,url:"http://www.baidu.com/a/b/c/redpackageaddress?id=100000","amount":300,"memo":"恭喜发财"}
6、MEDIA_TYPE_EMOJ 6
	{id:1,userid:2,dstid:3,cmd:10,media:6,"content":"cry"}
7、MEDIA_TYPE_Link 6
	{id:1,userid:2,dstid:3,cmd:10,media:7,"url":"http://www.a,com/dsturl.html"}
8、MEDIA_TYPE_VIDEO 8
	{id:1,userid:2,dstid:3,cmd:10,media:8,pic:"http://www.baidu.com/a/log,jpg",url:"http://www.a,com/a.mp4"}
9、MEDIA_TYPE_CONTACT 9
	{id:1,userid:2,dstid:3,cmd:10,media:9,"content":"10086","pic":"http://www.baidu.com/a/avatar,jpg","memo":"胡大力"}
*/

func init() {
	go udpSendProc()
	go udpRecvProc()
}

var udpsendchan chan []byte = make(chan []byte, 1024) //用来存放发送的要广播的数据

//形成userid和Node的映射关系
type Node struct {
	Conn      *websocket.Conn //并行转串行,
	DataQueue chan []byte
	GroupSets set.Interface
}

var ClientMap map[int64]*Node = make(map[int64]*Node, 0) //映射关系表
var RWLocker sync.RWMutex                                //读写锁

// ws发送协程
func (c *ChatService) SendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				global.Logger.Error(err.Error())
				return
			}
		}
	}
}

// ws接收协程
func (c *ChatService) RecvProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			global.Logger.Error(err.Error())
			return
		}
		dispatch(data)
		//把消息广播到局域网
		c.BroadMsg(data)
		global.Logger.Infof("[ws]<=%s\n", data)
	}
}

// 发送消息
func (c *ChatService) SendMsg(userId int64, msg []byte) {
	RWLocker.RLock()
	node, ok := ClientMap[userId]
	RWLocker.RUnlock()
	if ok {
		node.DataQueue <- msg
	}
}

//将消息广播到局域网
func (c *ChatService) BroadMsg(data []byte) {
	udpsendchan <- data
}

// udp数据的发送协程
func udpSendProc() {
	global.Logger.Infoln("start udpsendproc")
	// 使用udp协议拨号
	con, err := net.DialUDP("udp", nil,
		&net.UDPAddr{
			IP:   net.IPv4(192, 168, 0, 255),
			Port: 3000,
		})
	defer con.Close()
	if err != nil {
		global.Logger.Errorln(err.Error())
		return
	}

	for {
		select {
		case data := <-udpsendchan:
			_, err = con.Write(data)
			if err != nil {
				global.Logger.Errorln(err.Error())
				return
			}
		}
	}
}

// 完成upd接收并处理功能
func udpRecvProc() {
	global.Logger.Infoln("start udprecvproc")
	// 监听udp广播端口
	con, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 3000,
	})
	defer con.Close()
	if err != nil {
		global.Logger.Errorln(err.Error())
	}
	// 处理端口发过来的数据
	for {
		var buf [512]byte
		n, err := con.Read(buf[0:])
		if err != nil {
			global.Logger.Errorln(err.Error())
			return
		}
		//直接数据处理
		dispatch(buf[0:n])
	}
}

//后端调度逻辑处理
func dispatch(data []byte) {
	// 解析data为message
	msg := Message{}
	c := ChatService{}

	err := json.Unmarshal(data, &msg)
	if err != nil {
		global.Logger.Errorln(err.Error())
		return
	}

	// 根据cmd对逻辑进行处理
	switch msg.Cmd {
	case CMD_SINGLE_MSG:
		c.SendMsg(msg.Dstid, data)
	case CMD_ROOM_MSG:
		// 群聊转发逻辑
		for userId, v := range ClientMap {
			if v.GroupSets.Has(msg.Dstid) {
				//自己排除,不发送
				if msg.Userid != userId {
					v.DataQueue <- data
				}

			}
		}
	case CMD_HEART:
		// 一般啥都不做
	}
}
