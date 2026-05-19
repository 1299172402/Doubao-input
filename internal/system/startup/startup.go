package startup

import (
	"fmt"
	"runtime"
)

// InstallStartup 安装开机自启
func InstallStartup() error {
	switch runtime.GOOS {
	case "windows":
		return installWindowsStartup()
	case "darwin":
		return installMacOSStartup()
	case "linux":
		return installLinuxStartup()
	default:
		return fmt.Errorf("此系统不支持开机自启功能")
	}
}

// UninstallStartup 卸载开机自启
func UninstallStartup() error {
	switch runtime.GOOS {
	case "windows":
		return uninstallWindowsStartup()
	case "darwin":
		return uninstallMacOSStartup()
	case "linux":
		return uninstallLinuxStartup()
	default:
		return fmt.Errorf("此系统不支持开机自启功能")
	}
}

// UpdateStartup 更新开机自启状态
func UpdateStartup(enabled bool) error {
	if enabled {
		if err := InstallStartup(); err != nil {
			return fmt.Errorf("安装开机自启失败: %w", err)
		}
	} else {
		if err := UninstallStartup(); err != nil {
			return fmt.Errorf("卸载开机自启失败: %w", err)
		}
	}
	return nil
}
