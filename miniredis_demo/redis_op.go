// redis_op.go
package miniredis_demo

import (
	"context"
	"strings"
	"time"

	"github.com/redis/go-redis/v9" // 导入Go语言的Redis客户端库（版本8）
)

// 常量定义：有效网站列表的Redis键名
const (
	KeyValidWebsite = "app:valid:website:list"
)

// DoSomethingWithRedis 对Redis进行一系列操作的函数
// 参数：
//
//	rdb - Redis客户端实例
//	key - 要操作的Redis键
//
// 返回：
//
//	操作是否成功的布尔值
func DoSomethingWithRedis(rdb *redis.Client, key string) bool {
	// 创建上下文，用于控制Redis操作的生命周期
	ctx := context.TODO()

	// 检查key是否存在于有效网站集合中（KeyValidWebsite是一个Set类型）
	// SIsMember方法返回key是否存在于集合中的布尔值
	if !rdb.SIsMember(ctx, KeyValidWebsite, key).Val() {
		return false // 若不存在，直接返回失败
	}

	// 从Redis中获取key对应的值
	// Get方法返回值和可能的错误
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return false // 获取失败，返回失败
	}

	// 检查值是否以"https://"开头
	// 若不是，则添加该前缀
	if !strings.HasPrefix(val, "https://") {
		val = "https://" + val
	}

	// 将处理后的值存入名为"blog"的键中，并设置5秒过期时间
	// Set方法接收键名、值和过期时间（0表示永不过期）
	if err := rdb.Set(ctx, "blog", val, 5*time.Second).Err(); err != nil {
		return false // 设置失败，返回失败
	}

	return true // 所有操作成功，返回成功
}
