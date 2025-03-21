package main

import(
	"log"
)
func main() {
		//Missing code to receive auth token - THE AUTH TOKEN CAN NOT BE PUBLIC
		authToken = ""
		sender = ""
        //retrieveMessage(sender,authToken)
		//sendMessage(sender,authToken)
		triggerMessage()
		
		//For simple copy & pasting until I make an official UI:

		//488929085381410816 - Gage
		//488909657482592257 - Mordred
		//488905706804740126 - Sean
		//810039358530912256 - Bot
}
//error catcher (For redundancy)
func errc(err error) {
	if err != nil {
		log.Fatal(err)
	}
}