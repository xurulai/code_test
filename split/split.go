package split

// Split 函数用于将字符串 s 按分隔符 sep 进行分割
func Split(s, sep string) []string {
    var result []string
    sepLen := len(sep)
    if sepLen == 0 {
        return []string{s} // 空分隔符处理
    }
    
    start := 0
    for {
        idx := findSubstring(s, sep, start)
        if idx == -1 {
            break
        }
        result = append(result, s[start:idx])
        start = idx + sepLen
    }
    result = append(result, s[start:])
    return result
}

// findSubstring 在字符串 s 中从 start 位置开始查找子串 substr
// 返回子串首次出现的索引，如果未找到则返回 -1
func findSubstring(s, substr string, start int) int {
    for i := start; i <= len(s)-len(substr); i++ {
        if s[i:i+len(substr)] == substr {
            return i
        }
    }
    return -1
}