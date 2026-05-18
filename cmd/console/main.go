package main

import (
	"fmt"

	"Doubao-input/info"
	"Doubao-input/internal/core"
)

func main() {
	// 交互式模式（双击启动）
	fmt.Printf("Doubao Input\n")
	fmt.Printf("Version: %s\n", info.Version)

	// 启动消息监听
	core.StartClipboardWriter()
}
