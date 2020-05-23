/**
* @Author: kiosk
* @Mail: weijiaxiang007@foxmail.com
* @Date: 2020/3/28
**/
package template

import "time"

const (
	COMMUNITY_CATE_COM = 0x01
)

// 群信息
type Community struct {
	Id         int64     `orm:"pk auto bigint(20)" json:"id"`
	Gruopname  string    `orm:"varchar(30)" json:"group_name" description:"群名称"`
	Ownerid    int64     `orm:"bigint(20)" json:"ownerid" description:"群主ID"`
	Icon       string    `orm:"varchar(250)" json:"icon" description:"群logo"`
	Cate       int       `orm:"int(11)" json:"cate" description:"类型"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)" json:"update_time" description:"更新时间"`
	CreatTime  time.Time `orm:"auto_now_add;type(datetime)" json:"creat_time" description:"创建时间"`
}

// 设置引擎为 INNODB
func (u *Community) TableEngine() string {
	return "INNODB"
}
