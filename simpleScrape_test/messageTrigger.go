package main

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"log"
	"time"
	_ "github.com/mattn/go-sqlite3"
)

const (
	//Database file location - NOTE: make this generic
	dbPath   = "C:/Users/DJ/AppData/Local/Microsoft/Windows/Notifications/wpndatabase.db"
	//Updates every second
	interval = 1 * time.Second
	//Time to give after calling too many times
	cooldown = 10 * time.Second
)
//Payload data format:
//<toast><visual><binding template="ToastImageAndText02"><image id="1" src="C:\Users\username\AppData\Local\Temp\scoped_dir10692_939240952\sender.png"/><text id="1">sender_name</text><text id="2">sender_text</text></binding></visual><audio silent="true"/></toast>
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
func processClock() {
	for {
		var x int = 0
		for x < 100{
			sender, senderText := processDatabase(dbPath)
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
		return "",""
	}
	defer db.Close()

	rows, err := db.Query("SELECT Payload FROM Notification") // Replace with your actual table name
	if err != nil {
		log.Printf("Error querying database: %v", err)
		return "",""
	}
	defer rows.Close()

	var payloads []string
	for rows.Next() {
		var payload string
		if err := rows.Scan(&payload); err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}
		payloads = append(payloads, payload)
	}
	//First entry in payload is junk data so we ignore it.
	if len(payloads) <= 1 {
		log.Println("Not enough entries in the table. Skipping processing.")
		return "",""
	}
	secondPayload := payloads[1]
	var parsedPayload PayloadData
	if err := xml.Unmarshal([]byte(secondPayload), &parsedPayload); err != nil {
		log.Printf("Error parsing XML: %v", err)
		return "",""
	}
	fmt.Println(parsedPayload.Visual.Binding.Text[0].Text +": "+parsedPayload.Visual.Binding.Text[1].Text+"\n")
	return parsedPayload.Visual.Binding.Text[0].Text, parsedPayload.Visual.Binding.Text[1].Text
}