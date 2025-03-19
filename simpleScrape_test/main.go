package main

import (
        "encoding/json"
        "fmt"
		"io"
        "net/http"
		"log"
)

func retrieveMessages(channelID string) string{
        url := "https://discord.com/api/v9/channels/" + channelID + "/messages"
        req, err := http.NewRequest("GET", url, nil)
		errc(err)

        req.Header.Set("Authorization", "MzQ1NjM2ODMzODg0ODk3Mjkx.GM5UM9.hzvSLNncqaX19i_wo1kKxeB60PSxKOc4jfyLwI")

        resp, err := http.DefaultClient.Do(req)
		errc(err)

        defer resp.Body.Close()

        body, err := io.ReadAll(resp.Body)
		errc(err)

        var jsonn []map[string]interface{}
        err = json.Unmarshal(body, &jsonn)
		errc(err)

		message := jsonn[0]["content"].(string)
		uname := jsonn[0]["author"].(map[string]interface{})["username"].(string)
		fmt.Println(uname)
		if(uname != "boommag"){
			fmt.Println(message)
			return uname + ": " + message
		}
		return ""
}

func main() {
        retrieveMessages("488929085381410816")
}
func errc(err error) {
	if err != nil {
		log.Fatal(err)
	}
}