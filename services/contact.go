/**
* @Author: kiosk
* @Mail: weijiaxiang007@foxmail.com
* @Date: 2020/5/16
**/
package services

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"imchat-api/models"
	"imchat-api/models/template"
)

type ContactArg struct {
	PageArg
	Userid int64 `json:"userid" form:"userid"`
	Dstid  int64 `json:"dstid"  form:"dstid"`
}

type ContactService struct {
}

// 搜寻好友
func (c *ContactService) SearchFriend(userId int64) (userList []UserInfoJson, err error) {
	db := models.GenerateDBOrm()

	contacts := make([]template.Contact, 0)
	users := make([]template.User, 0)
	userList = make([]UserInfoJson, 0)
	objIds := make([]int64, 0)

	contact := new(template.Contact)
	if _, err = db.QueryTable(contact).
		Filter("ownerid", userId).
		Filter("cate", template.CONCAT_CATE_USER).All(&contacts); err != nil {
		return userList, err
	}

	for _, v := range contacts {
		objIds = append(objIds, v.Dstobj)
	}

	if len(objIds) == 0 {
		return userList, err
	}

	user := new(template.User)
	if _, err = db.QueryTable(user).Filter("id__in", objIds).All(&users); err != nil {
		return userList, err
	}
	for _, u := range users {
		userInfo := UserInfoJson{
			Id:       u.Id,
			Nickname: u.Nickname,
			Mobile:   u.Mobile,
			Avatar:   u.Avatar,
			Token:    "",
			Sex:      u.Sex,
			Memo:     u.Memo,
		}
		userList = append(userList, userInfo)
	}
	return userList, err
}

// 添加好友
func (c *ContactService) AddFriend(userId int64, dstId int64) error {
	db := models.GenerateDBOrm()
	if userId == dstId {
		return errors.New("不能添加自己为好友")
	}
	//判断是否已经加了好友
	contact := new(template.Contact)
	//查询是否已经是好友
	if exist := db.QueryTable(contact).
		Filter("ownerid", userId).
		Filter("dstobj", dstId).
		Filter("cate", template.CONCAT_CATE_USER).
		Exist(); exist {
		return errors.New("该用户已添加过好友")
	}

	if err := db.Begin(); err != nil {
		return errors.New("数据库事物失败")
	}

	myContact := template.Contact{
		Ownerid: userId,
		Dstobj:  dstId,
		Cate:    template.CONCAT_CATE_USER,
		Memo:    "",
	}
	_, err1 := db.Insert(&myContact)

	dstContact := template.Contact{
		Ownerid: dstId,
		Dstobj:  userId,
		Cate:    template.CONCAT_CATE_USER,
		Memo:    "",
	}
	_, err2 := db.Insert(&dstContact)
	if err1 == nil && err2 == nil {
		_ = db.Commit()
		return nil
	} else {
		_ = db.Rollback()
		return errors.New("插入数据库失败")
	}
}

// 搜寻群
func (c *ContactService) SearchCommunity(userId int64) (communityList []template.Community, err error) {
	db := models.GenerateDBOrm()

	contacts := make([]template.Contact, 0)
	communityList = make([]template.Community, 0)
	comIds := make([]int64, 0)

	contact := new(template.Contact)
	if _, err = db.QueryTable(contact).
		Filter("ownerid", userId).
		Filter("cate", template.CONCAT_CATE_COMUNITY).All(&contacts); err != nil {
		return communityList, err
	}

	for _, v := range contacts {
		comIds = append(comIds, v.Dstobj)
	}

	if len(comIds) == 0 {
		return communityList, err
	}

	community := new(template.Community)
	if _, err = db.QueryTable(community).Filter("id__in", comIds).All(&communityList); err != nil {
		return communityList, err
	}

	return communityList, err
}

// 搜寻群
func (c *ContactService) SearchCommunityIds(userId int64) (commIds []int64) {
	db := models.GenerateDBOrm()

	comIds := make([]int64, 0)
	var maps []orm.Params
	if _, err := db.QueryTable(new(template.Contact)).
		Filter("ownerid", userId).
		Filter("cate", template.CONCAT_CATE_COMUNITY).
		Values(&maps, "id"); err == nil {
		for _, m := range maps {
			comIds = append(comIds, m["Id"].(int64))
		}
	}
	return comIds
}

// 创建群
func (c *ContactService) CreateCommunity(comm template.Community) (template.Community, error) {
	db := models.GenerateDBOrm()

	if len(comm.Gruopname) == 0 {
		return template.Community{},errors.New("缺少群名称")
	}
	if comm.Ownerid == 0 {
		return template.Community{},errors.New("请先登录")
	}

	cnt, err := db.QueryTable(new(template.Community)).Filter("ownerid", comm.Ownerid).Count()
	if err != nil {
		return template.Community{},err
	}
	_ = db.Begin()

	if cnt >= 5 {
		return template.Community{},errors.New("一个用户最多只能创见5个群")
	}else {
		_, err1 := db.Insert(&comm)

		contact := &template.Contact{
			Ownerid:    comm.Ownerid,
			Dstobj:     comm.Id,
			Cate:       template.CONCAT_CATE_COMUNITY,
			Memo:       "",
		}

		_, err2 := db.Insert(&contact)

		if err1 != nil || err2 != nil {
			db.Rollback()
			return template.Community{}, errors.New("插入数据库失败")
		}else {
			_ = db.Commit()
			return comm, nil
		}
	}
}



