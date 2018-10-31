package controllers

import (
	"github.com/astaxie/beego"
	"zerogo/utils/Constant"
	"zerogo/utils"
	"log"
	"zerogo/common/manager/RedisManager"
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

	tvTopList := RedisManager.GetVideosByKeyAndTag(Constant.VIDEO_PREFIX_HOME_TV_HOT_KEY, Constant.CRAWLER_TYPE_LETV);
	liveList := RedisManager.GetVideosByKeyAndTag(Constant.VIDEO_PREFIX_HOME_LIVE_KEY, Constant.CRAWLER_TYPE_PANDA);
	c.Data["CarouselPic"] = carousePicList
	c.Data["Recommend"] = recommendList
	c.Data["TV"] = tvList
	c.Data["Cartoon"] = cartoonList
	c.Data["Movie"] = lvMovieList
	c.Data["TVTop"] = tvTopList
	c.Data["Live"] = liveList
}

func (c *MainController) SourceDetail() {
	c.InitDefault()
	sourceUrl := utils.DecodeUrl(c.GetString("u"))
	log.Println("url:", sourceUrl)
	c.Data["Source"] = ""
	c.TplName = "videoDetailTest.html"
}

func (c *MainController) See() {
	c.InitDefault()

}

func (c *MainController)InitDefault() {
	c.Data["Website"] = Constant.WEB_SITE
	c.Data["Email"] = Constant.EMAIL
	c.Data["EmailName"] = Constant.AUTHOR
}
