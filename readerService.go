package main

import (
	"errors"
	"time"
)

type ReaderService struct {
	Config    *Config
	AuthToken string
}

func NewReaderService(cfg *Config) *ReaderService {
	return &ReaderService{
		Config:    cfg,
		AuthToken: "test_token", // 这里可以设置一个默认的 token
	}
}

func (r *ReaderService) StartReading() error {
	// 1. 检查授权
	lic, err := LoadLicense()
	if err != nil || lic.ValidTill.Before(time.Now()) {
		return errors.New("授权已过期")
	}

	// 2. 启动模拟器（ADB 启动）
	if err := StartEmulator(); err != nil {
		return err
	}

	// 3. 打开浏览器访问目标网站
	if err := OpenBrowserInEmulator("https://jd.com"); err != nil {
		return err
	}

	// 4. 登录（三方 API，或 Cookie 注入）
	if r.AuthToken == "" {
		return errors.New("未登录")
	} else {
		//执行登录
	}

	// 5. 自动阅读文章流程
	// 这里可以添加自动阅读的逻辑，比如模拟点击、滚动等操作

	return nil
}
