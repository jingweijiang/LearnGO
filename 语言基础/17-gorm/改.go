package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main4() {

	dataSourceName := "root:123456@tcp(127.0.0.1:3306)/tester?charset=utf8&parseTime=True&loc=Local"
	client, err := gorm.Open(mysql.Open(dataSourceName), nil)
	checkErr(err)

	// 修改一条数据

	//client.Model(&User{}).Where("id=?", 3).Update("city", "日本")

	// 修改多条数据
	client.Model(&User{}).Where("id=?", 3).Updates(
		map[string]interface{}{"city": "美国", "gender": "男"})

}
