package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func sendMessage(channelID string, authToken string) {
    payload := map[string]string{
        "content": "TEST2",
    }

    url := "https://discord.com/api/v9/channels/"+ channelID +"/messages"

    client := &http.Client{}
    req, err := http.NewRequest("POST", url, mapToForm(payload))
	errc(err)

    req.Header.Set("Authorization",authToken)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    resp, err := client.Do(req)
	errc(err)
    defer resp.Body.Close()

    body, _ := io.ReadAll(resp.Body)

    fmt.Println(string(body))
}

func mapToForm(m map[string]string) *strings.Reader {
    b := strings.Builder{}
    for k, v := range m {
        fmt.Fprintf(&b, "%s=%s&", k, v)
    }
    return strings.NewReader(b.String()[:len(b.String())-1])
}