package utils

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var redisClient *redis.Client

func SetUpRedis() {
	data, err := os.ReadFile("./config/config.yml")
	if err != nil {
		log.Fatalf("error reading config file: %v", err)
	}

	var config Configuration
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error unmarshaling config: %v", err)
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Address, // Redis服务器地址
		Password: "",                   // 密码
		DB:       0,                    // 使用的数据库索引
	})

}
func GetRedisClient() *redis.Client {
	if DB == nil {
		panic("can not connect redis")
	}
	return redisClient
}
func SetExpiredToken(key string, val string) error {

	client := GetRedisClient()
	ctx := context.Background()

	// 执行Redis命令 token过期时间是24h
	err := client.Set(ctx, "token:"+key, val, 60*60*24).Err()
	return err
}

func IsIncludeExpiredToken(key string) bool {
	client := GetRedisClient()
	ctx := context.Background()

	exists, err := client.Exists(ctx, key).Result()
	if err != nil {
		fmt.Println("Failed to check key existence:", err)
		return false
	}

	if exists == 1 {
		return true
	} else {
		return false
	}
}
