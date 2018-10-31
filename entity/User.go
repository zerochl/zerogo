package entity

import "time"

type User struct {
	Id        int64

	// 用户唯一身份识别 ID
	OpenId    string

	// 密码（暂时用不到）
	Password  string

	/**
	 * 登录类型 {@link LoginType}
	 */
	LoginType string

	// 昵称
	NickName  string `orm:"column(nickname)"`

	// 头像
	Avatar    string

	// 性别
	Gender    string

	// 其他信息
	Meta      string

	// 用户信息 MD5 值，用于校验用户信息是否休息
	Md5       string

	CreateTime time.Time
}
