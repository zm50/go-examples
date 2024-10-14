package main

import (
	"context"
	"fmt"
	"go-examples/grpc/hello"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 连接远程服务器
	conn, err := grpc.NewClient("localhost:40404", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 创建客户端实例
	cli := hello.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 调用服务端方法
	resp, err := cli.SayHello(ctx, &hello.HelloRequest{Name: "world"})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Greeting: %s\n", resp.Message)
}
