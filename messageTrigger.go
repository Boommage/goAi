package main

import (
	"database/sql"
	"fmt"
	_"github.com/mattn/go-sqlite3"
	"time"
)

var sender string;

func clock() {
	db, err := sql.Open("sqlite3", "./messages.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Loop forever, checking for new messages
	for {
		rows, err := db.Query("SELECT id, username, message FROM messages WHERE processed = 0")
		if err != nil {
			panic(err)
		}
		_, err = db.Exec("PRAGMA journal_mode=WAL;")
		if err != nil {
    		panic(err)
		}

		var id int
		var username, message string

		for rows.Next() {
			err = rows.Scan(&id, &username, &message)
			if err != nil {
				continue
			}

			// Process the message
			//fmt.Printf("Go received: %s -> %s\n", username, message)
			sender = username
			retrieveMessage(username, message)
			time.Sleep(1000)
			_, err = db.Exec("DELETE FROM messages WHERE id = ?", id)
			if err != nil {
    		fmt.Println("Error deleting message:", err)
			}
		}
		rows.Close()

		time.Sleep(2 * time.Second)
	}
}