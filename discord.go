package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type webhook struct {
	Content string `json:"content"`
}

// TODO: 可変長引数に対応し第二引数にwebhookのURLを受け取るようにする
// TODO: 環境変数か第二引数でwebhookのURLを取得し、取得できない場合はエラーを返すようにする
func SendMessage(message string) error {
	webhook := webhook{
		Content: message,
	}

	jsonData, err := json.Marshal(webhook)
	if err != nil {
		return err
	}

	if os.Getenv("DISCORD_WEBHOOK") == "" {
		return fmt.Errorf("DISCORD_WEBHOOK is not set")
	}

	req, err := http.NewRequest("POST", os.Getenv("DISCORD_WEBHOOK"), bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return err
	}

	return nil
}
