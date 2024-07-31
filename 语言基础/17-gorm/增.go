package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main2() {

	dataSourceName := "root:123456@tcp(127.0.0.1:3306)/tester?charset=utf8&parseTime=True&loc=Local"
	client, err := gorm.Open(mysql.Open(dataSourceName), nil)
	checkErr(err)

	// 新增一条数据
	user := User{
		Uid:     1001,
		Keyword: "张三",
		City:    "上海",
	}
	err = client.Create(&user).Error
	//client.CreateInBatches()
	checkErr(err)
	fmt.Printf("新增一条数据: %v\n", user)

}
