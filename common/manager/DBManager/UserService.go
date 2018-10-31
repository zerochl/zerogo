package DBManager

import (
	"zerogo/entity"
	"zerogo/db/Dao"
)

type UserService interface {
	/**
	     * 登录时，更新用户信息
	     */
	UpdateUserInfo(user entity.User) entity.User
	/**
	     * 获取用户信息
	     * @param userId 用户Id
	     * @return 用户信息
	     */
	GetUserInfo(userId int64) entity.User
	GetUserInfoByOpenId(openId string)entity.User
	/**
	     * 新用户
	     */
	GetNewUsers(size int) []entity.User
	/**
	     * 活跃用户，一周内，收藏内容越多，排名越高
	     */
	GetActiveUsers(size int) []entity.User
	/**
	     * 人气用户, 粉丝越多，排名越高
	     */
	GetPopularUsers(size int) []entity.User
	/**
	     * 获取用户粉丝
	     */
	GetFans(userId int64, page int) []entity.User
	/**
	     * 获取用户偶像
	     */
	GetIdols(userId int64, page int) []entity.User
}

var userDao Dao.UserDao

//func (c *UserService)UpdateUserInfo(user entity.User) entity.User  {
//	openId := user.OpenId
//	avatar := user.Avatar
//	avatar = strings.Replace(avatar,"http:","",-1)
//	user.Avatar = avatar
//	origin :=
//}
//
//func (c *UserService) GetUserInfoByOpenId(openId string)entity.User{
//	var err error
//	var user entity.User
//}