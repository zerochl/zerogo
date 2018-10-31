package entity

import "AnyVideo-Go/utils/Constant"

type Video struct {
	Id         int64
	/* 视频名称 */
	Title      string

	/* 视频图片 */
	Image      string

	/* 视频播放地址 */
	PlayUrl    string

	/* 播放类型 */
	PlayType   string

	/* [版权] 视频源地址 */
	Value      string

	/* [版权] 视频提供方 */
	Provider   string

	/* [版权] 视频解析方名称 */
	ParserName string

	/* [版权] 视频解析方官网 */
	Parser     string

	/* 其他信息 */
	Other      string
}

func NewVideo(provider string) *Video {
	video := &Video{}
	video.Provider = provider
	video.ParserName = Constant.AUTHOR
	video.Parser = Constant.AUTHOR_WEBSITE
	return video
}