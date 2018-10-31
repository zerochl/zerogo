package controllers

import (
	"github.com/astaxie/beego"
	"zerogo/utils"
	"log"
)

type ApiController struct {
	beego.Controller
}

func (c *ApiController) GetVideoParseInfo(){
	sourceUrl := utils.DecodeUrl(c.GetString("v"))
	log.Println("url:",sourceUrl)
	//source := ParserManager.Parse(sourceUrl)
	c.Data["json"] = ""
	c.ServeJSON()
}

func (c *ApiController)GetEpisodes()  {
	sourceUrl := utils.DecodeUrl(c.GetString("v"))
	log.Println("in url:",sourceUrl)
	//videoParse := ParserManager.GetParseByUrl(sourceUrl)
	//source := videoParse.ParseEpisodes(sourceUrl)
	//source := ParserManager.ParseEpisodes(sourceUrl)
	c.Data["json"] = ""
	c.ServeJSON()
}