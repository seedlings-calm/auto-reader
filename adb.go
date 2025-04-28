package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

type ADBService struct {
}

func (a *ADBService) ADBShell(command string) (string, error) {
	out, err := exec.Command("adb", "shell", command).Output()
	return string(out), err
}
func checkADB() bool {
	_, err := os.Stat("C:\\Program Files\\雷电模拟器\\adb.exe")
	return !os.IsNotExist(err)
}
func StartEmulator() error {
	// 检查雷神模拟器是否已经安装
	if !checkADB() {
		fmt.Println("雷神模拟器未安装，正在下载并安装...")
		err := downloadAndInstall()
		if err != nil {
			log.Fatalf("模拟器安装失败: %v", err)
		}
		fmt.Println("模拟器安装成功")
	}
	cmd := exec.Command("adb", "start-server")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to start ADB server: %w", err)
	}
	return nil
}

func OpenBrowserInEmulator(url string) error {
	cmd := exec.Command("adb", "shell", "am", "start", "-a", "android.intent.action.VIEW", "-d", url)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to OpenBrowserInEmulator url: %w", err)
	}
	return nil
}

// 模拟器下载链接
const downloadURL = "https://example.com/雷神模拟器/ThunderEmulatorSetup.exe"

// 安装文件的保存路径
const installPath = "C:\\temp\\ThunderEmulatorSetup.exe"

func downloadAndInstall() error {
	// 下载模拟器
	out, err := os.Create(installPath)
	if err != nil {
		return fmt.Errorf("failed to create file for download: %w", err)
	}
	defer out.Close()

	resp, err := http.Get(downloadURL)
	if err != nil {
		return fmt.Errorf("failed to download ADB setup: %w", err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save downloaded file: %w", err)
	}

	// 下载完成后，自动运行安装程序
	cmd := exec.Command(installPath, "/S") // "/S" 参数通常表示静默安装
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to install the emulator: %w", err)
	}

	// 安装后，检查是否成功
	if !checkADB() {
		return fmt.Errorf("simulator installation failed")
	}

	return nil
}
