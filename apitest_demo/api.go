package apitest_demo

// api.go

import (
	"bytes"         // 用于处理字节切片
	"encoding/json" // JSON编解码
	"io/ioutil"     // 输入输出工具（Go 1.16+建议使用os和io包替代）
	"net/http"      // HTTP客户端和服务器实现
)

// ReqParam 定义发送给API的请求参数结构
// JSON标签指定了字段在JSON数据中的名称
type ReqParam struct {
	X int `json:"x"`
}

// Result 定义API返回结果的结构
// JSON标签指定了字段在JSON数据中的名称
type Result struct {
	Value int `json:"value"`
}

// GetResultByAPI 通过调用外部API获取结果并进行后续处理
// 参数:
//
//	x - 发送给API的请求参数
//	y - 用于后续计算的本地参数
//
// 返回:
//
//	API返回值加上本地参数y的结果，或错误时返回-1
func GetResultByAPI(x, y int) int {
	// 创建请求参数对象并序列化为JSON字节切片
	// 忽略错误处理（这是代码的一个问题点）
	p := &ReqParam{X: x}
	b, _ := json.Marshal(p)

	// 发送HTTP POST请求到指定API
	// 参数:
	//   URL - API端点
	//   Content-Type - 指定请求体格式为JSON
	//   Request Body - 请求参数的JSON字节表示
	resp, err := http.Post(
		"http://your-api.com/post",
		"application/json",
		bytes.NewBuffer(b),
	)

	// 错误处理：如果请求失败，返回-1表示错误
	// 这里没有记录错误信息，不利于调试
	if err != nil {
		return -1
	}

	// 确保响应体在函数结束时关闭，防止资源泄漏
	defer resp.Body.Close()

	// 读取响应体的全部内容
	// 忽略错误处理（这是代码的另一个问题点）
	body, _ := ioutil.ReadAll(resp.Body)

	// 将响应体JSON数据反序列化为Result结构体
	var ret Result
	if err := json.Unmarshal(body, &ret); err != nil {
		return -1
	}

	// 对API返回的数据做后续处理（这里是加上本地参数y）
	return ret.Value + y
}
