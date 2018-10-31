package controllers

import (
	"github.com/astaxie/beego"
	"AnyVideo-Go/service/manager/RedisManager"
	"AnyVideo-Go/utils/Constant"
	"AnyVideo-Go/service/tool"
	"AnyVideo-Go/service/manager/ParserManager"
	"AnyVideo-Go/entity"
	"AnyVideo-Go/utils"
	"log"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.InitDefault()
	c.TplName = "index.html"
	carousePicList := RedisManager.GetVideosByKeyAndTag(Constant.VIDEO_PREFIX_HOME_CAROUSEL_KEY, Constant.CRAWLER_TYPE_LETV);
	recommendList := RedisManager.GetVideosByKeyAndTag(Constant.VIDEO_PREFIX_HOME_RECOMMEND_KEY, Constant.CRAWLER_TYPE_LETV);
	tvList := RedisManager.GetVideosByKeyAndTag(Constant.VIDEO_PREFIX_HOME_TV_KEY, Constant.CRAWLER_TYPE_LETV);
	cartoonList := RedisManager.GetVideosByKeyAndTag(Constant.VIDEO_PREFIX_HOME_CARTOON_KEY, Constant.CRAWLER_TYPE_LETV);

	lvMovieList := RedisManager.GetVideosByKeyAndTag(Constant.VIDEO_PREFIX_HOME_MOVIE_KEY, Constant.CRAWLER_TYPE_LETV);
	qqMovieList := RedisManager.GetVideosByKeyAndTag(Constant.VIDEO_PREFIX_HOME_MOVIE_KEY, Constant.CRAWLER_TYPE_QQ);

	tvTopList := RedisManager.GetVideosByKeyAndTag(Constant.VIDEO_PREFIX_HOME_TV_HOT_KEY, Constant.CRAWLER_TYPE_LETV);
	liveList := RedisManager.GetVideosByKeyAndTag(Constant.VIDEO_PREFIX_HOME_LIVE_KEY, Constant.CRAWLER_TYPE_PANDA);
	c.Data["CarouselPic"] = carousePicList
	c.Data["Recommend"] = recommendList
	c.Data["TV"] = tvList
	c.Data["Cartoon"] = cartoonList
	c.Data["Movie"] = tool.IntervalBlend(lvMovieList, qqMovieList)
	c.Data["TVTop"] = tvTopList
	c.Data["Live"] = liveList
}

func (c *MainController) SourceDetail() {
	c.InitDefault()
	sourceUrl := utils.DecodeUrl(c.GetString("u"))
	log.Println("url:", sourceUrl)
	source := ParserManager.Parse(sourceUrl)
	c.Data["Source"] = source
	if _, ok := source.(*entity.Video); ok {
		c.TplName = "videoDetailTest.html"
	} else {
		c.TplName = "articleDetail.html"
	}
}

func (c *MainController) See() {
	c.InitDefault()

}

func (c *MainController)InitDefault() {
	c.Data["Website"] = Constant.WEB_SITE
	c.Data["Email"] = Constant.EMAIL
	c.Data["EmailName"] = Constant.AUTHOR
}
