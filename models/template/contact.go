/**
* @Author: kiosk
* @Mail: weijiaxiang007@foxmail.com
* @Date: 2020/3/28
**/
package template

import "time"

const (
	CONCAT_CATE_USER     = 0x01
	CONCAT_CATE_COMUNITY = 0x02
)

//好友和群都存在这个表里面
type Contact struct {
	Id         int64     `orm:"pk auto bigint(20)" json:"id" `
	Ownerid    int64     `orm:"bigint(20)" json:"ownerid" description:"本端id"`
	Dstobj     int64     `orm:"bigint(20)" json:"dstobj" description:"对端信息"`
	Cate       int       `orm:"int(11)" json:"cate" description:"类型"`
	Memo       string    `orm:"varchar(120)" json:"memo" description:"备注"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)" json:"update_time" description:"更新时间"`
	CreatTime  time.Time `orm:"auto_now_add;type(datetime)" json:"creat_time" description:"创建时间"`
}

// 设置引擎为 INNODB
func (u *Contact) TableEngine() string {
	return "INNODB"
}
