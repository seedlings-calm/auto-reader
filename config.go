package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	ReadTime   int    `json:"read_time"` // 单篇文章秒数
	SessionKey string `json:"session_key"`
}

type ConfigService struct {
}

func (c *ConfigService) GetConfig() (*Config, error) {
	data, err := os.ReadFile("config.json")
	if err != nil {
		return nil, err
	}
	var cfg Config
	_ = json.Unmarshal(data, &cfg)
	return &cfg, nil
}

func (c *ConfigService) SetConfig(cfg *Config) error {
	data, _ := json.Marshal(cfg)
	return os.WriteFile("config.json", data, 0644)
}
