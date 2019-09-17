package controllers

import (
	"context"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/astaxie/beego"
	"google.golang.org/grpc"

	"github.com/hieuvecto/todo-grpc/pkg/api/v1"

)

type CreateController struct {
	beego.Controller
}

func (c *CreateController) Post() {

	address := beego.AppConfig.String("grpc-server")
	apiVersion := beego.AppConfig.String("apiVersion")

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect to grpc server: %v", err)
	}
	defer conn.Close()

	client := v1.NewToDoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	t := time.Now().In(time.UTC)
	insert_at, _ := ptypes.TimestampProto(t)
	// Call Create
	req := v1.CreateRequest{
		Api: apiVersion,
		ToDo: &v1.ToDo{
			Title:       c.GetString("title"),
			Description: c.GetString("description"),
			InsertAt:    insert_at,
			UpdateAt: 	insert_at,
		},
	}
	res, err := client.Create(ctx, &req)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", res)
	
	c.Redirect("/", 302)
}
