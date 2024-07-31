package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func redisString(ctx context.Context, client *redis.Client) {

	key := "mykey"
	value := "大脸猫"

	// 设置， 可以设置过期时间， 0 表示永不过期
	//err := client.Set(ctx, key, value, 0).Err()
	err := client.Set(ctx, key, value, 1*time.Second).Err()
	if err != nil {
		panic(err)
	}

	// 设置过期时间
	client.Expire(ctx, key, 3*time.Second)

	// 获取
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	// 删除
	err = client.Del(ctx, key).Err()
	if err != nil {
		panic(err)
	}
}

func redisList(ctx context.Context, client *redis.Client) {
	key := "mylist"
	values := []interface{}{1, 2, 3, "apple", "banana", "orange"}
	// 左侧插入元素
	//client.LPush(ctx, key, values...)
	// 右侧插入元素
	err := client.RPush(ctx, key, values...).Err() //values...表示可变参数
	if err != nil {
		panic(err)
	}

	val, err := client.LRange(ctx, key, 0, 100).Result()
	fmt.Println(val, err)

	client.Del(ctx, key)
}

func redisHash(ctx context.Context, client *redis.Client) {
	client.HSet(ctx, "1001", "name", "大脸猫", "age", 18, "gender", "男")
	client.HSet(ctx, "1002", "name", "小怪宏", "age", 20, "gender", "女")

	// 获取学生1的name和age
	val, err := client.HMGet(ctx, "学生1", "name", "age").Result()
	fmt.Println(val, err)

	// 获取所有学生的name和age
	val1, err := client.HGetAll(ctx, "1001").Result()
	fmt.Println(val1, err)
	val2, err := client.HGetAll(ctx, "1002").Result()
	fmt.Println(val2, err)

	//client.Del(ctx, "学生1", "学生2")

}

func main() {
	fmt.Println("Redis 操作")
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	var ctx = context.TODO()

	fmt.Println("Pinging Redis server")
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Failed to connect to Redis:", err)
		panic(err)
	}
	//redisString(ctx, rdb)
	//redisList(ctx, rdb)
	redisHash(ctx, rdb)
}
