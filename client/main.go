package main

import (
	"github.com/astaxie/beego"
	_ "github.com/hieuvecto/todo-grpc/client/routers"
)

func main() {
	beego.Run()
}
