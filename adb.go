package main

import (
	"fmt"
	"os/exec"
)

type ADBService struct {
}

func (a *ADBService) ADBShell(command string) (string, error) {
	out, err := exec.Command("adb", "shell", command).Output()
	return string(out), err
}

func StartEmulator() error {
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
