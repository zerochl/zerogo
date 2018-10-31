package main

import (
	_ "AnyVideo-Go/routers"
	_ "AnyVideo-Go/initial"
	"github.com/astaxie/beego"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	beego.Run()
}
