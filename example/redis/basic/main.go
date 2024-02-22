package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/nico612/go-project/pkg/db"
	"github.com/redis/go-redis/v9"
	"time"
)

func main() {
	redis := db.GetRedisOr(&db.Options{
		Addrs:       []string{"127.0.0.1:6379"},
		DB:          0,
		DialTimeout: 10 * time.Second,
		ReadTimeout: 10 * time.Second,
		PoolSize:    50,
	})

	stringOperate(redis)
}

// 字符串操作
func stringOperate(rdb redis.UniversalClient) {

	ctx := context.Background()

	// 设置值, 仅当 key 不存在时
	rdb.SetNX(ctx, "username", "zhangsan", 0)

	username, err := rdb.Get(ctx, "username").Result()
	if err != nil {
		fmt.Println("get username failed, err:", err)
		return
	}
	fmt.Println("username:", username)

	// 获取不存在健的值
	_, err = rdb.Get(ctx, "user").Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			fmt.Println("key does not exist")
			return
		}
		fmt.Println("get username failed, err:", err)
		return
	}

}
