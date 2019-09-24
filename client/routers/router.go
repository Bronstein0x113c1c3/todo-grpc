package routers

import (
	"github.com/astaxie/beego"
	"github.com/hieuvecto/todo-grpc/client/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/todo", &controllers.TodoController{}, "post:CreateTodo")
	beego.Router("/todo/:id:int", &controllers.TodoController{}, "get:ReadTodo;delete:DeleteTodo")
	beego.Router("/todo/:id:int/edit", &controllers.TodoController{}, "post:EditTodo")
}
