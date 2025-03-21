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

// URL is optional, if not provided it will use the DISCORD_WEBHOOK environment variable
// if that is not set it will return an error
func SendMessage(message string, url ...string) error {
	var webhookUrl string
	if len(url) > 0 {
		webhookUrl = url[0]
	} else {
		webhookUrl = os.Getenv("DISCORD_WEBHOOK")
	}

	if webhookUrl == "" {
		return fmt.Errorf("DISCORD_WEBHOOK is not set")
	}

	webhook := webhook{
		Content: message,
	}

	jsonData, err := json.Marshal(webhook)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", webhookUrl, bytes.NewBuffer(jsonData))
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
