package routers

import (
	"zerogo/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/view",&controllers.MainController{},"get:SourceDetail")

	beego.Router("/api/episode",&controllers.ApiController{},"get:GetEpisodes")
	beego.Router("/api/video",&controllers.ApiController{},"get:GetVideoParseInfo")

	var FilterMethod = func(ctx *context.Context) {
		if ctx.Input.Query("_method") != "" && ctx.Input.IsPost() {
			ctx.Request.Method = ctx.Input.Query("_method")
		}
		ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	}
	beego.InsertFilter("*", beego.BeforeRouter, FilterMethod)
}
