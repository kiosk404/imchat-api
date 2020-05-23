/**
* @Author: kiosk
* @Mail: weijiaxiang007@foxmail.com
* @Date: 2020/3/28
**/
package template

import "time"

const (
	SEX_WOMEN  = "W"
	SEX_MEN    = "M"
	SEX_UNKNOW = "U"
)

type User struct {
	//用户ID
	Id       int64  `orm:"pk autoincr bigint(20)" json:"id"`
	Mobile   string `orm:"varchar(20)" json:"mobile" description:"手机号"`
	Password string `orm:"varchar(40)" json:"password" description:"密码"`
	Avatar   string `orm:"varchar(150)" json:"avatar" description:"身份"`
	Sex      string `orm:"varchar(2)"  json:"sex" description:"性别"`
	Nickname string `orm:"varchar(20)" json:"nickname" description:"昵称"`
	//加盐随机字符串6
	Salt   string `orm:"varchar(10)" json:"salt" description:"盐"`
	Online int    `orm:"int(10)" form:"online" json:"online" description:"是否在线"`
	//前端鉴权因子,
	Token      string    `orm:"varchar(40)" json:"token" description:"token"`
	Memo       string    `orm:"varchar(140)" form:"memo" json:"memo" description:"备注"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)" json:"update_time" description:"上次更新时间"`
	CreatTime  time.Time `orm:"auto_now_add;type(datetime)" json:"creat_time" description:"创建时间"`
}

// 设置引擎为 INNODB
func (u *User) TableEngine() string {
	return "INNODB"
}
