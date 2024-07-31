package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type User struct {
	ID      int
	Uid     int
	Keyword string `gorm:"column:keywords"`
	City    string
}

func (User) TableName() string {
	return "user"
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)

		os.Exit(1)
	}
}

func readAll(client *gorm.DB, city string) []User {
	var users []User
	err := client.Select("id", "uid", "keywords", "city", "degree").Where("city =?", city).Find(&users).Error
	checkErr(err)
	return users
}

func readOne(client *gorm.DB, city string) *User {
	var users User
	//users.ID = 2	// 这里可以设置ID，也可以不设置,会和where条件一起使用，and
	err := client.Select("id", "uid", "keywords", "city", "degree").Where("city =?", city).Find(&users).Error
	checkErr(err)
	return &users
}

func mai1() {

	dataSourceName := "root:123456@tcp(127.0.0.1:3306)/tester?charset=utf8&parseTime=True&loc=Local"
	client, err := gorm.Open(mysql.Open(dataSourceName), nil)
	checkErr(err)

	// 查找一条数据
	usersOne := readOne(client, "上海")
	fmt.Printf("%+v", usersOne)

	// 查找所有数据
	//usersAll := readAll(client, "上海")
	//fmt.Println(usersAll)

	//if len(users) > 0 {
	//	for _, user := range users {
	//		fmt.Printf("%+v", user)
	//	}
	//} else {
	//	fmt.Println("No user found")
	//}

}
