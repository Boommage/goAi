package main
//import "fmt"

//All users will be stored here
var users = map[string]string{
	//"averytinymedic" : "488929085381410816",
	//"murderousreptile" : "534515514832060436",
	//"destroyed spines" : "488905706804740126",
	//"fny" : "488908433249140749",
	//"vic !" : "1182059215922143325",
	"djevilevil" : "1411819459668082854",
	//"disbelief" : "488909657482592257",
}
//currently does not work...
var userNames = map[string]string{
	"Thomas" : "fny",
	"Sean" : "destroyed spines",
	"May" : "djevilevil",
}
//Updated id based off the current new sender
var currentID string = ""

//Checks if the user exists, send message if so
func retrieveMessage(user string, message string){
	print(user," ",message)
	if id, exists := users[user]; exists{
		currentID = id
		runAI(message)
	}
}