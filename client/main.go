package main

import (
	"github.com/astaxie/beego"
	func_map "github.com/hieuvecto/todo-grpc/client/func_map"
	_ "github.com/hieuvecto/todo-grpc/client/routers"
)

func main() {
	beego.AddFuncMap("mod", func_map.Mod)
	beego.AddFuncMap("timeElapsed", func_map.TimeElapsed)
	beego.AddFuncMap("timeElapsedTimestamp", func_map.TimeElapsedTimestamp)
	beego.Run()
}
