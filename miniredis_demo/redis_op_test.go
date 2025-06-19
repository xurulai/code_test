// redis_op_test.go

package miniredis_demo

import (
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2" // 内存中的Redis模拟器
	"github.com/redis/go-redis/v9"     // Redis客户端库
)

// TestDoSomethingWithRedis 测试DoSomethingWithRedis函数的行为
func TestDoSomethingWithRedis(t *testing.T) {
	// 创建并启动内存中的Redis服务器
	s, err := miniredis.Run()
	if err != nil {
		panic(err) // 启动失败时终止测试
	}
	defer s.Close() // 测试结束后关闭服务器

	// 准备测试数据：
	// 1. 设置键"q1mi"的值为"liwenzhou.com"
	s.Set("q1mi", "liwenzhou.com")
	// 2. 将"q1mi"添加到有效网站集合中
	s.SAdd(KeyValidWebsite, "q1mi")

	// 创建Redis客户端，连接到内存中的Redis服务器
	rdb := redis.NewClient(&redis.Options{
		Addr: s.Addr(), // 使用miniredis分配的地址
	})

	// 调用被测试函数，传入模拟的Redis客户端和测试键
	ok := DoSomethingWithRedis(rdb, "q1mi")
	if !ok {
		t.Fatal("函数执行失败") // 如果函数返回false，测试失败
	}

	// 验证结果：手动检查
	// 获取"blog"键的值，并验证是否为预期的"https://liwenzhou.com"
	got, err := s.Get("blog")
	if err != nil || got != "https://liwenzhou.com" {
		t.Fatalf("'blog'键的值不符合预期")
	}

	// 验证结果：使用miniredis提供的辅助函数
	// 功能同上，但使用miniredis的断言方法
	s.CheckGet(t, "blog", "https://liwenzhou.com")

	// 测试过期时间：
	// 1. 将时间快进5秒（模拟5秒后）
	s.FastForward(5 * time.Second)
	// 2. 检查"blog"键是否已过期
	if s.Exists("blog") {
		t.Fatal("'blog'键应该已过期")
	}
}
