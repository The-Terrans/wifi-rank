package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

type result struct {
	Time     time.Time `json:"time"`
	SSID     string    `json:"ssid"`
	BSSID    string    `json:"bssid"`
	Download float64   `json:"download_mbps"`
	Upload   float64   `json:"upload_mbps"`
	PingMs   float64   `json:"ping_ms"`
	Server   string    `json:"server"`
}

// historyPath is ~/.config/wifi-rank/history.jsonl on Linux.
func historyPath() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "wifi-rank", "history.jsonl"), nil
}

func loadHistory(path string) ([]result, error) {
	f, err := os.Open(path)
	if errors.Is(err, fs.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var results []result
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		var r result
		if err := json.Unmarshal(sc.Bytes(), &r); err != nil {
			return nil, err
		}
		results = append(results, r)
	}
	return results, sc.Err()
}

func appendHistory(path string, r result) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(r)
}
