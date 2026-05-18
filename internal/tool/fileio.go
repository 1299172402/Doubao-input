package tool

import (
	"fmt"
	"os"
	"strings"
)

// ReadCurlFile 从文件读取 curl 命令
func ReadCurlFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("请检查 session.txt 文件是否存在: %w", err)
	}
	// 处理行尾反斜杠续行
	content := strings.ReplaceAll(string(data), "\\\n", " ")
	return content, nil
}

// WriteCurlFile 将 curl 命令写入文件
func WriteCurlFile(path string, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}
