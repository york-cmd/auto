package utils

import (
	"auto/commands"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func SendMessageToServerChan(title, message string) {
	if commands.Config.ServerSendKey == "" {
		return
	}
	serverChanURL := fmt.Sprintf("https://sctapi.ftqq.com/%s.send", commands.Config.ServerSendKey)
	data := url.Values{}
	data.Set("text", title)
	data.Set("desp", message)
	resp, err := http.PostForm(serverChanURL, data)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("SendMessageToServerChan error : %v", err)
	}
}
