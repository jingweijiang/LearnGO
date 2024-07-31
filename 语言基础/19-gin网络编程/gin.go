package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"strconv"
)

type Student struct {
	Name   string
	Age    int
	Gender string
}

func GetStudentInfo(studentId string) Student {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	var ctx = context.TODO()

	stu := Student{}
	val, _ := rdb.HGetAll(ctx, studentId).Result()
	stu.Name = val["name"]
	stu.Age, _ = strconv.Atoi(val["age"])
	stu.Gender, _ = val["gender"]

	return stu
}

func GetName(ctx *gin.Context) {
	studentId := ctx.Query("studentId")
	if len(studentId) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "studentId is required"})
	}
	stu := GetStudentInfo(studentId)
	ctx.JSON(http.StatusOK, gin.H{"name": stu.Name})
	return
}

func GetAge(ctx *gin.Context) {
	studentId := ctx.PostForm("studentId")
	if len(studentId) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "studentId is required"})
	}
	stu := GetStudentInfo(studentId)
	ctx.JSON(http.StatusOK, gin.H{"age": stu.Age})
	return
}

type Request struct {
	StudentId string `form:"studentId" json:"studentId" xml:"studentId"  binding:"required"`
}

func GetGender(ctx *gin.Context) {

	var param Request
	if err := ctx.ShouldBind(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	if len(param.StudentId) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "studentId is required"})
	}
	stu := GetStudentInfo(param.StudentId)
	ctx.JSON(http.StatusOK, gin.H{"gender": stu.Gender})
	return
}

func main() {
	router := gin.Default()

	router.GET("/get_name", GetName)
	router.POST("/get_age", GetAge)
	router.POST("/get_gender", GetGender)

	err := router.Run("0.0.0.0:2345")
	if err != nil {
		panic(err)
	}
}
