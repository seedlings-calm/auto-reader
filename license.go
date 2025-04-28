package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Auth struct {
}

type License struct {
	// The name of the license
	Key       string    `json:"key"`
	ValidTill time.Time `json:"valid_till"`
}

func (a *Auth) VerifyKey(key string) bool {
	if IsValidKey(key) {
		_ = SaveLicense(key)
		return true
	}
	return false
}

func (a *Auth) CheckLicense() bool {
	_, err := LoadLicense()
	return err == nil
}

func (a *Auth) LoginAndGetTicket(username, password string) (string, error) {
	form := url.Values{}
	form.Add("username", username)
	form.Add("password", password)

	req, _ := http.NewRequest("POST", "https://example.com/login", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 假设票据在 Set-Cookie 或者响应中
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "session_id" {
			return cookie.Value, nil
		}
	}
	return "", errors.New("no session_id found")
}

func IsValidKey(key string) bool {
	return strings.HasPrefix(key, "VIP-") && len(key) > 10
}

func SaveLicense(key string) error {
	lic := License{
		Key:       key,
		ValidTill: time.Now().Add(30 * 24 * time.Hour),
	}
	data, _ := json.Marshal(lic)
	return os.WriteFile("license.json", data, 0644)
}

func LoadLicense() (*License, error) {
	data, err := os.ReadFile("license.json")
	if err != nil {
		return nil, err
	}
	var lic License
	_ = json.Unmarshal(data, &lic)
	if time.Now().After(lic.ValidTill) {
		return nil, errors.New("license expired")
	}
	return &lic, nil
}
