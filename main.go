package main

import(
	"log"
)
//Missing code to receive auth token - THE AUTH TOKEN CAN NOT BE PUBLIC
var authToken string = ""
//The person sending the dm
//var sender string = ""
func main() {
		//Calls the message reader when a notification is received 
		//testAI()
		clock()
}
//error catcher (For redundancy)
func errc(err error) {
	if err != nil {
		log.Fatal(err)
	}
}