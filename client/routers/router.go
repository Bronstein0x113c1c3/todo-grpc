package routers

import (
	"github.com/astaxie/beego"
	"github.com/hieuvecto/todo-grpc/client/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
