// Triggers when the user receives a dm.
package main

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	//All windows notifications are sent to a database file.
	//This is the files location 
	// - NOTE: make this generic
	dbPath   = "C:/Users/djgam/AppData/Local/Microsoft/Windows/Notifications/wpndatabase.db"
	//Reads from db file every second
	interval = 1 * time.Second
	//10 second cooldown from reading the db file.	
	cooldown = 10 * time.Second
)

var sender string
var senderText string
//Every notif db file has a "payload"
/*
---Payload data format---

//<toast>
// 	<visual>
// 		<binding template="ToastImageAndText02">
// 			<image id="1" src="C:\Users\username\AppData\Local\Temp\scoped_dir\sender.png"/>
// 			<text id="1">sender_name</text>
// 			<text id="2">sender_text</text>
// 		</binding>
// 	</visual>
// 	<audio silent="true"/>
//</toast>
*/
type PayloadData struct {
	XMLName xml.Name `xml:"toast"`
	Visual  struct {
		Version string `xml:"version,attr"`
		Binding struct {
			Text []TextElement `xml:"text"`
		} `xml:"binding"`
	} `xml:"visual"`
}
//The most important part of the payload, the sender name and sender text
type TextElement struct {
	ID   string `xml:"id,attr"`
	Text string `xml:",chardata"`
}

//Reads from the db file on an interval.
func clock() {
	for {
		var x int = 0
		for x < 100{
			senderTemp, senderTextTemp := processDatabase(dbPath)
			//If the message text and sender are the same then don't recall	
			if(senderTemp != sender && senderTextTemp != senderText){
				sender = strings.ToLower(senderTemp)
				senderText = senderTextTemp
				retrieveMessage(sender,senderText)
			}
			time.Sleep(interval)
			x++
		}
		time.Sleep(cooldown)
	}
}
//Extracts data from db file
func processDatabase(dbFile string) (sender string, senderText string){
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Printf("Error opening database: %v", err)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT Payload FROM Notification") // Replace with your actual table name
	if err != nil {
		log.Printf("Error querying database: %v", err)
		return
	}
	defer rows.Close()

	var toastPayload string
	for rows.Next() {
		var payload string
		if err := rows.Scan(&payload); err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}
		if strings.HasPrefix(strings.TrimSpace(payload), "<toast>") {
			toastPayload = payload
			break // Stop at the first <toast> entry found
		}
	}

	if toastPayload == "" {
		log.Println("No <toast> entry found.")
		return
	}

	var parsedToast PayloadData
	if err := xml.Unmarshal([]byte(toastPayload), &parsedToast); err != nil {
		log.Printf("Error parsing XML: %v", err)
		return
	}

	fmt.Println("First <toast> entry found:")
	fmt.Println(parsedToast.Visual.Binding.Text[0].Text +": "+parsedToast.Visual.Binding.Text[1].Text+"\n")
	return parsedToast.Visual.Binding.Text[0].Text, parsedToast.Visual.Binding.Text[1].Text
}