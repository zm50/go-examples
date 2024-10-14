package main

import (
	"context"
	"fmt"
	"go-examples/grpc/hello"
	"net"

	"google.golang.org/grpc"
)

// 接口实现
type GreeterServer struct {
	hello.UnimplementedGreeterServer
}

// 实现SayHello方法
func (s *GreeterServer) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloReply, error) {
	fmt.Printf("Received: %s\n", in.GetName())

	return &hello.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	// 创建监听器
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 40404))
	if err != nil {
		panic(err)
	}
	defer lis.Close()

	// 创建grpc server
	s := grpc.NewServer()

	// 注册服务
	hello.RegisterGreeterServer(s, &GreeterServer{})

	// 启动服务
	s.Serve(lis)
}
