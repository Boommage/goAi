package main

import (
        "encoding/json"
        "fmt"
		"io"
        "net/http"
)

//Retrieves the first message of a given user's discord dm 
func retrieveMessage(channelID string, authToken string) string{

	//Obtains discord link off channel id
	url := "https://discord.com/api/v9/channels/" + channelID + "/messages"

	temp := ""

	if( )

	//Creates a HTTP GET request directly to the dm
    req, err := http.NewRequest("GET", url, nil)
	errc(err)

	//Sends discord API authorization to Go's default HTTP client
	req.Header.Set("Authorization", authToken)
	//Receives the servers response
	resp, err := http.DefaultClient.Do(req)
	errc(err)
	//Releases the network resource
	defer resp.Body.Close()

	//Stores the response as a byte slice for easy reading
	body, err := io.ReadAll(resp.Body)
	errc(err)

	//Extracts data contents from the body as a map
	//Entries 0-49 each contain individual message data
    var jsonn []map[string]interface{}
    err = json.Unmarshal(body, &jsonn)
	errc(err)

	//Grabs the first message string - AKA entry 0
	 message := jsonn[0]["content"].(string)

	//Grabs the user of the first message
	 uname := jsonn[0]["author"].(map[string]interface{})["username"].(string)

	//If the user who sent the message is not me(boommag, hi :| ) 
	//Then return the message to the AI
	 fmt.Println(uname)
	 if(uname != "boommag"){
	 	fmt.Println(message)
	 	return uname + ": " + message
	 }
	//else return nothing
	return ""
}