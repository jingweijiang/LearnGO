package main

import (
	"context"
	"encoding/json"
	"fmt"
	myproto "g6/idl/my_proto"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"log"
	"net"
)

type StudentServer struct {
}

type Student struct {
	Name      string             `json:"name"`
	Age       int32              `json:"age"`
	Gender    bool               `json:"gender"`
	Height    float64            `json:"height"`
	Locations []string           `json:"locations"`
	Scores    map[string]float32 `json:"scores"`
}

// 从 Redis 中获取学生信息
func GetStudentInfo(studentId string) myproto.Student {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer rdb.Close()
	var ctx = context.TODO()

	stu := myproto.Student{}
	val, _ := rdb.Get(ctx, studentId).Result()

	err := json.Unmarshal([]byte(val), &stu)
	if err != nil {
		fmt.Println("Failed to unmarshal JSON data:", err)
		return myproto.Student{}
	}

	return stu
}

// 实现 StudentService 的rpc接口
func (s *StudentServer) GetStudentInfo(ctx context.Context, req *myproto.StudentRequest) (*myproto.Student, error) {

	defer func() {
		if r := recover(); r != nil {
			log.Printf("panic: %v", r)
		}
	}()
	var stu myproto.Student

	stu = GetStudentInfo(req.StudentId)
	return &stu, nil
}

func main() {
	// 启动 grpc 服务
	listen, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		panic(err)

	}

	server := grpc.NewServer()
	// 注册 StudentService
	myproto.RegisterStudentServiceServer(server, &StudentServer{})

	// 启动 grpc 服务
	if err := server.Serve(listen); err != nil {
		panic(err)
	}

}
