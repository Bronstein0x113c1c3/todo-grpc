package controllers

import (
	"context"
	"log"
	"time"

	"github.com/astaxie/beego"

	"google.golang.org/grpc"

	v1 "github.com/hieuvecto/todo-grpc/pkg/api/v1"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	address := beego.AppConfig.String("grpc-server")
	apiVersion := beego.AppConfig.String("apiVersion")

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect to grpc server: %v", err)
		c.Abort("500")
	}
	defer conn.Close()

	client := v1.NewToDoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := v1.ReadAllRequest{
		Api: apiVersion,
	}
	res, err := client.ReadAll(ctx, &req)
	if err != nil {
		log.Fatalf("ReadAll failed: %v", err)
		c.Abort("500")
	}

	c.Data["ToDos"] = res.ToDos

	c.TplName = "index.tpl"
}
