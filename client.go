package main

import (
	"fmt"

	pb "org.springbus/api"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	Address = "127.0.0.1:50052"
)

func main() {

	conn, err := grpc.Dial("127.0.0.1:1234", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	//初始化客户端
	c := pb.NewHelloClient(conn)

	req := new(pb.HelloReq)
	req.Req = "hello"
	req.GoodsName = "xx口罩"
	r, err := c.SayHello(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}

	m := pb.NewMsgSvrClient(conn)
	mr, _ := m.DoSvr(context.Background(), req)

	fmt.Println(r.Rep)
	fmt.Println(mr.Rep)

}
