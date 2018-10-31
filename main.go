package main

import (
	_ "zerogo/routers"
	_ "zerogo/initial"
	"github.com/astaxie/beego"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	beego.Run()
}
