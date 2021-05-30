package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"

	pb "mygithub/cat_grpc_demo/proto/hello"
	"net"
)

const Address string = "127.0.0.1:50052"

type helloSserver struct {
	// todo 这个要继承
	pb.HelloServiceServer
}

//var HelloServer = server{}

// SayHello 实现Hello服务接口
func (h *helloSserver) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.", in.Name)
	return resp, nil
}

//func (h *server) mustEmbedUnimplementedHelloServiceServer()  {
//
//}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &helloSserver{})
	grpclog.Info("Listen on " + Address)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
