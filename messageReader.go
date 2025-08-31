package main
//import "fmt"

//All users will be stored here
var users = map[string]string{
	//"averytinymedic" : "488929085381410816",
	//"murderousreptile" : "534515514832060436",
	"destroyed spines" : "488905706804740126",
	"fny" : "488908433249140749",
	//"vic !" : "1182059215922143325",
	"boommag " : "757127014112034867",
	//"disbelief" : "488909657482592257",
}
var userNames = map[string]string{
	"Thomas" : "fny",
	"Sean" : "destroyed spines",
}
//Updated id based off the current new sender
var currentID string = ""

//Checks if the user exists, send message if so
func retrieveMessage(sender string, text string){
	if id, exists := users[sender]; exists{
		currentID = id
		if userNames[sender] != "" {
			sender = userNames[sender]
		}
		runAI(text)
	}
}