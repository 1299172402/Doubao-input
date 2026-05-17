package main

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"fyne.io/systray"
)

//go:embed static/logo.png
var logo2PNG []byte

func StartTray() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(logo2PNG)
	systray.SetTitle("Doubao Input")
	systray.SetTooltip("Doubao Input " + Version)
	mOpen := systray.AddMenuItem("打开 Web 界面", "在浏览器中打开管理界面")
	mClose := systray.AddMenuItem("关闭 Web 界面", "关闭 Web 服务器")
	mClose.Hide()
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("退出", "退出程序")

	go func() {
		webRunning := false
		for {
			select {
			case <-mOpen.ClickedCh:
				if !webRunning {
					// 启动 Web 服务
					addr := ":2828"
					if p := os.Getenv("DOUBAO_INPUT_PORT"); p != "" {
						addr = ":" + p
					}
					go StartWeb(addr)
					fmt.Printf("Web 界面: http://localhost%s\n", addr)
					webRunning = true
					mClose.Show()
				}
				openBrowser("http://localhost:2828")
			case <-mClose.ClickedCh:
				// 关闭 Web 服务
				StopWeb()
				webRunning = false
				fmt.Println("Web 服务已关闭")
				mClose.Hide()
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
	// clean up here
}

func openBrowser(url string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}
	cmd.Start()
}
