package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	siteUpCheckTimeInterval = 10 * time.Second
	siteUrl                 = "https://google.com"
	webhookURL              = "https://discord.com/api/webhooks/WEBHOOK_ID/WEBHOOK_TOKEN"
)

func main() {
	for range time.Tick(siteUpCheckTimeInterval) {
		resp, err := http.Get(siteUrl)
		if err != nil {
			log.Fatal(err)
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Println("Website Response Is:", resp.StatusCode, http.StatusText(resp.StatusCode))
			sendWebhook("Website is down!")
		} else {
			fmt.Println("All Good Chef!")
		}

		resp.Body.Close()
	}
}

func sendWebhook(message string) {
	payload := `{"content":"` + message + `"}`

	resp, err := http.Post(webhookURL, "application/json", strings.NewReader(payload))
	if err != nil {
		log.Println("Failed to send webhook:", err)
	} else {
		defer resp.Body.Close()
		fmt.Println("Webhook notification sent successfully")
	}
}
