package controllers

import (
	"context"
	"log"
	"time"
	"strconv"

	"github.com/golang/protobuf/ptypes"
	"github.com/astaxie/beego"
	"google.golang.org/grpc"

	"github.com/hieuvecto/todo-grpc/pkg/api/v1"

)

type TodoController struct {
	beego.Controller
}

func (c *TodoController) ReadTodo() {

	address := beego.AppConfig.String("grpc-server")
	apiVersion := beego.AppConfig.String("apiVersion")
	id := c.Ctx.Input.Param(":id")

	id_int, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.Abort("500")
	}

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

	req := v1.ReadRequest{
		Api: apiVersion,
		Id:  id_int,
	}
	res, err := client.Read(ctx, &req)
	if err != nil {
		log.Fatalf("Read failed: %v", err)
		c.Abort("500")
	}
	log.Printf("Read result: <%+v>\n\n", res)

	c.Data["ToDo"] = res.ToDo

	c.TplName = "read.tpl"
}

func (c *TodoController) CreateTodo() {

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
		c.Abort("500")
	}
	log.Printf("Create result: <%+v>\n\n", res)
	
	c.Redirect("/", 302)
}

func (c *TodoController) EditTodo() {
	address := beego.AppConfig.String("grpc-server")
	apiVersion := beego.AppConfig.String("apiVersion")
	id := c.Ctx.Input.Param(":id")

	id_int, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.Abort("500")
	}

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

	// Update
	t := time.Now().In(time.UTC)
	update_at, _ := ptypes.TimestampProto(t)

	req := v1.UpdateRequest{
		Api: apiVersion,
		ToDo: &v1.ToDo{
			Id:          id_int,
			Title:       c.GetString("title"),
			Description: c.GetString("description"),
			UpdateAt:    update_at,
		},
	}

	res, err := client.Update(ctx, &req)
	if err != nil {
		log.Fatalf("Update failed: %v", err)
		c.Abort("500")
	}
	log.Printf("Update result: <%+v>\n\n", res)
	
	c.Redirect("/", 302)
}

func (c *TodoController) DeleteTodo() {
	address := beego.AppConfig.String("grpc-server")
	apiVersion := beego.AppConfig.String("apiVersion")
	id := c.Ctx.Input.Param(":id")

	id_int, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.Abort("500")
	}

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

	req := v1.DeleteRequest{
		Api: apiVersion,
		Id:  id_int,
	}
	res, err := client.Delete(ctx, &req)
	if err != nil {
		log.Fatalf("Delete failed: %v", err)
		c.Abort("500")
	}
	log.Printf("Delete result: <%+v>\n\n", res)
	
	c.ServeJSON()
}