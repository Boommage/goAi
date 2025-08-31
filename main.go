package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)
//load env
var err = godotenv.Load()
//Missing code to receive auth token for every user 
var authToken string = os.Getenv("AUTHTOKEN")

//The person sending the dm
//var sender string = ""
func main() {
		//testAI()
		//Calls the message reader when a notification is received 
		clock()
}
//error catcher (For redundancy)
func errc(err error) {
	if err != nil {
		log.Fatal(err)
	}
}