package main

import (
	"context"
	myproto "g6/idl/my_proto"
	"google.golang.org/grpc"
	"log"
	"testing"
	"time"
)

// 调用GetStudentInfo rpc服务测试函数
func TestGetStudentInfo(t *testing.T) {
	// 连接grpc服务端
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 关闭连接
	defer conn.Close()

	// 创建客户端实例
	c := myproto.NewStudentServiceClient(conn)
	// WithTimeout设置超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 调用rpc服务
	r, err := c.GetStudentInfo(ctx, &myproto.StudentRequest{StudentId: "student:1001"})
	if err != nil {
		log.Fatalf("could not get student info: %v", err)
	}
	log.Printf("Student Info: %v", r)

}
