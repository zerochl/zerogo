package Dao

import (
	"AnyVideo-Go/entity"
	"github.com/astaxie/beego"
	"strconv"
	"github.com/astaxie/beego/orm"
	"time"
	"errors"
	"zerogo/utils/Exception"
)

const (
	tableName = "any_user"
	nickName = "nickname"
)

type UserDao interface {
	Insert(user entity.User) int
	Update(user entity.User)
	SelectByOpenId(openId string) entity.User
	SelectById(userId int64) entity.User
	SelectNew(size int) []entity.User
	SelectActive(size int) []entity.User
	SelectPopular(size int) []entity.User
	SelectIdols(userId int64,begin int,end int) []entity.User
	SelectFans(userId int64,begin int,end int) []entity.User
}

func (c *UserDao) Insert(user entity.User) int{
	orm := orm.NewOrm()
	err := c.CheckNewUser(user)
	if nil != err {
		return err
	}
	user.CreateTime = time.Now()
	id, err := orm.Insert(user)
	if nil == err {
		return err
	}
	user.Id = id
	return nil
}

func (c *UserDao) SelectByOpenId(openId string) entity.User{
	orm := orm.NewOrm()
}

func (c *UserDao)CheckNewUser(user *entity.User) error {
	name := user.NickName
	pass := user.Password

	err := c.CheckUserPsw(pass)
	if nil != err {
		return err
	}

	err = c.CheckUserName(name)
	if nil != err {
		return err
	}

	return nil
}

func (c *UserDao)CheckUserPsw(pass string) error{
	minPassLength := 4
	maxPassLength := 10
	passLength := len(pass)

	if passLength < minPassLength || passLength > maxPassLength {
		return errors.New("密码长度只能在" + strconv.Itoa(minPassLength) + "-" + strconv.Itoa(maxPassLength) + "字符之间")
	}
	return nil
}

func (c *UserDao)CheckUserName(name string) error {

	minNameLength := 4
	maxNameLength := 10

	beego.Error("minNameLength:", minNameLength, "maxNameLength:", maxNameLength)
	nameLength := len(name)
	if nameLength < minNameLength || nameLength > maxNameLength {
		return errors.New("用户名长度只能在" + strconv.Itoa(minNameLength) + "-" + strconv.Itoa(maxNameLength) + "字符之间")
	}

	orm := orm.NewOrm()
	count, err := orm.QueryTable(tableName).Filter(nickName, name).Count()

	if nil != err || count > 0 {
		return Exception.USER_NAME_EXISTENT
	}
	return nil

}
