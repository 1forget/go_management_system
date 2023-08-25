package utils

import (
	"context"
	"fmt"
)

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
