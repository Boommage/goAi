package main

//import "fmt"

//All users will be stored here
var users = map[string]string{
	"averytinymedic" : "488929085381410816",
	//"murderousreptile" : "534515514832060436",
	"destroyed spines" : "488905706804740126",
	"fny" : "488908433249140749",
	//"disbelief" : "488909657482592257",

}
//Updated id based off the current new sender
var currentID string = ""

func retrieveMessage(sender string, text string){
	if id, exists := users[sender]; exists{
		currentID = id
		runAI(sender+": "+text)
	}
}