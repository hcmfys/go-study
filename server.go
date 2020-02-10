package main

import (
	"github.com/giorgisio/goav/avformat"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "org.springbus/api"
)

//定义一个helloServer并实现约定的接口
type helloService struct{}

func (h helloService) SayHello(ctx context.Context, in *pb.HelloReq) (*pb.HelloRep, error) {
	resp := new(pb.HelloRep)
	resp.Rep = "reply : " + in.Req + " :goodsName:" + in.GoodsName
	return resp, nil
}

func (h helloService) DoSvr(_ context.Context, in *pb.HelloReq) (*pb.HelloRep, error) {
	resp := new(pb.HelloRep)
	resp.Rep = "do msg : " + in.Req + " : goodsName:" + in.GoodsName
	return resp, nil
}

func readMp4File() {

	filename := "sample.mp4"

	// Register all formats and codecs
	avformat.AvRegisterAll()

	ctx := avformat.AvformatAllocContext()

	// Open video file
	if avformat.AvformatOpenInput(&ctx, filename, nil, nil) != 0 {
		log.Println("Error: Couldn't open file.")
		return
	}

	// Retrieve stream information
	if ctx.AvformatFindStreamInfo(nil) < 0 {
		log.Println("Error: Couldn't find stream information.")

		// Close input file and free context
		ctx.AvformatCloseInput()
		return
	}

}

var HelloServer = helloService{}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("failed to listen")
	}
	//实现gRPC Server
	serv := grpc.NewServer()
	//注册helloServer为客户端提供服务
	pb.RegisterHelloServer(serv, HelloServer) //内部调用了s.RegisterServer()
	pb.RegisterMsgSvrServer(serv, HelloServer)
	serv.Serve(listen)

}
