package httptest_demo

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_helloHandler(t *testing.T) {
	tests := []struct {
		name   string // 测试用例名称
		param  string // POST请求的JSON格式参数
		expect string // 期望返回的消息内容
	}{
		{"base case", `{"name": "liwenzhou"}`, "hello liwenzhou"}, // 正常情况
		{"bad case", "", "we need a name"},                        // 异常情况：缺少必要参数
	}

	r := SetupRouter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(
				"POST",                      // 请求方法
				"/hello",                    // 请求URL路径
				strings.NewReader(tt.param), // 请求体，将参数转为可读的数据流
			)

			// 创建一个响应记录器，用于捕获处理请求后的响应
			w := httptest.NewRecorder()

			// 使用Gin引擎处理模拟请求，并将响应写入记录器
			r.ServeHTTP(w, req)

			// 断言响应状态码是否为200 OK
			assert.Equal(t, http.StatusOK, w.Code)

			// 解析响应体JSON数据到map结构
			var resp map[string]string
			err := json.Unmarshal([]byte(w.Body.String()), &resp)

			// 断言JSON解析过程中没有错误
			assert.Nil(t, err)

			// 断言响应中的"msg"字段值是否符合预期
			assert.Equal(t, tt.expect, resp["msg"])
		})
	}
}
