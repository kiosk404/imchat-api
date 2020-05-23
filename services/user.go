/**
* @Author: kiosk
* @Mail: weijiaxiang007@foxmail.com
* @Date: 2020/5/15
**/
package services

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"imchat-api/models"
	"imchat-api/models/template"
	"imchat-api/utils"
	"math/rand"
	"time"
)

type UserService struct {
}

type UserInfoJson struct {
	Id       int64  `json:"id"`
	Nickname string `json:"nickname"`
	Mobile   string `json:"mobile"`
	Token    string `json:"token"`
	Sex      string `json:"sex"`
	Memo     string `json:"memo"`
	Avatar   string `json:"avatar"`
}

//注册函数
func (s *UserService) Register(user *template.User) (err error) {
	db := models.GenerateDBOrm()

	if name, empty := utils.CheckEmpty(&user.Mobile,
		&user.Password,
		&user.Nickname,
		&user.Memo,
		&user.Sex); !empty {
		return errors.New("参数" + name + "不能为空")
	}

	//检测手机号码是否存在,
	tmp := template.User{}

	exist := db.QueryTable(new(template.User)).Filter("mobile", user.Mobile).Exist()

	//如果存在则返回提示已经注册
	if exist {
		return errors.New("mobile :" + user.Mobile + ", 该手机号已经注册")
	}

	//否则拼接插入数据
	tmp.Mobile = user.Mobile
	tmp.Avatar = user.Avatar
	tmp.Nickname = user.Nickname
	tmp.Sex = user.Sex
	tmp.Memo = user.Memo
	tmp.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	tmp.Password = utils.MakePasswd(user.Password, tmp.Salt)
	//token 可以是一个随机数
	tmp.Token = fmt.Sprintf("%08d", rand.Int31())

	//插入 InsertOne
	_, err = db.Insert(&tmp)

	return err
}

//登录函数
func (s *UserService) Login(mobile, plainpwd string) (user template.User, err error) {

	db := models.GenerateDBOrm()

	//首先通过手机号查询用户
	tmp := template.User{}
	err = db.QueryTable(new(template.User)).Filter("mobile", mobile).One(&tmp)
	if err == orm.ErrMultiRows || err == orm.ErrNoRows {
		// 多条的时候报错 或 没有找到记录
		return tmp, errors.New("该用户不存在")
	}

	//查询到了比对密码
	if !utils.ValidatePasswd(plainpwd, tmp.Salt, tmp.Password) {
		return tmp, errors.New("密码不正确")
	}
	//刷新token,安全
	token := utils.MD5Encode(fmt.Sprintf("%d", time.Now().Unix()))
	tmp.Token = token

	_, err = db.Update(&tmp)
	return tmp, err
}

//查找某个用户
func (s *UserService) Find(userId int64) (user template.User) {
	db := models.GenerateDBOrm()

	tmp := template.User{}
	_ = db.QueryTable(new(template.User)).Filter("id", userId).One(&tmp)
	return tmp
}

//加载用户信息
func (s *UserService) LoadUserInfo(userInfoJson *UserInfoJson, user template.User) {
	userInfoJson.Nickname = user.Nickname
	userInfoJson.Mobile = user.Mobile
	userInfoJson.Token = user.Token
	userInfoJson.Sex = user.Sex
	userInfoJson.Id = user.Id
	userInfoJson.Memo = user.Memo
	userInfoJson.Avatar = user.Avatar
}

//检查用户是否合法
func (s *UserService) CheckToken(userId int64, token string) bool {
	db := models.GenerateDBOrm()

	user := template.User{}
	_ = db.QueryTable(new(template.User)).Filter("id", userId).One(&user)
	return token == user.Token
}
