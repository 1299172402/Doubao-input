package main

import (
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	// 直接使用 robotgo 的 Type 函数输入文本
	for {
		robotgo.Type("Hello, World!")
		time.Sleep(1 * time.Second)
	}
}
