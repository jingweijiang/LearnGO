package main

import (
	"context"
	myproto "g6/idl/my_proto"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

type StudentServer struct {
}

type Student struct {
	Name   string
	Age    int
	Gender string
}

func GetStudentInfo(studentId string) myproto.Student {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer rdb.Close()
	var ctx = context.TODO()

	stu := myproto.Student{}
	val, _ := rdb.HGetAll(ctx, studentId).Result()
	stu.Name = val["name"]
	stu.Age, _ = strconv.Atoi(val["age"])
	stu.Gender, _ = val["gender"]

	return stu
}

func (s *StudentServer) GetStudentInfo(ctx context.Context, req *myproto.StudentRequest) (*myproto.Student, error) {

	defer func() {
		if r := recover(); r != nil {
			log.Printf("panic: %v", r)
		}
	}()
	var stu myproto.Student

	req.StudentId = "1001"
	stu := GetStudentInfo(req.StudentId)
	return &stu, nil
}

func main() {
	listen, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		panic(err)

	}
	server := grpc.NewServer()
	student_service.RegisterStudentServiceServer(server, &StudentServer{})
	if err := server.Serve(listen); err != nil {
		panic(err)
	}

}
