package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	pb "mygithub/cat_grpc_demo/proto/hello"
)

const Address string = "127.0.0.1:50052"

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()

	c := pb.NewHelloServiceClient(conn)
	req := &pb.HelloRequest{Name: "grpc:shf"}
	res, err := c.SayHello(context.Background(), req)
	if err != nil {
		grpclog.Fatalln(err)
	}
	fmt.Println("获取", res.Message)
}
