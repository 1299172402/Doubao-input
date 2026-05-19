package system

import (
	_ "embed"
	_ "image/png"

	"fmt"

	"Doubao-input/assets"
	"Doubao-input/internal/config"
	"Doubao-input/internal/tool"

	"github.com/energye/systray"
)

func StartTray() {
	systray.Run(onReady, onExit)
}

func openSetting() {
	tool.OpenBrowser()
}

var mAutoType *systray.MenuItem

func taggleAutoType() {
	cfg := config.GetConfig()
	cfg.AutoType = !cfg.AutoType
	config.SaveConfig(cfg)
	if config.GetConfig().AutoType {
		mAutoType.Check()
	} else {
		mAutoType.Uncheck()
	}
}

func onReady() {
	icoData, err := tool.PngToIco(assets.LogoPNG)
	if err != nil {
		fmt.Println("图标转换失败:", err)
		icoData = assets.LogoPNG
	}
	systray.SetIcon(icoData)
	systray.SetTitle("豆包语音输入")
	systray.SetTooltip("豆包语音输入")

	mAutoType = systray.AddMenuItemCheckbox("自动输入", "启用自动输入功能", config.GetConfig().AutoType)
	mAutoType.Click(taggleAutoType)
	mOpen := systray.AddMenuItem("设置", "打开设置页面")
	mOpen.Click(openSetting)
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("退出", "退出程序")
	mQuit.Click(systray.Quit)
}

func onExit() {
	// clean up here
}
