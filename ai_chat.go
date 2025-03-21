package main

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tmc/langchaingo/chains"
	//"github.com/tmc/langchaingo/llms"
	//"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/memory/sqlite3"
)

/*Greetings Everyone!

This is a work in progress project i'm working on.
The goal is for Deepseek and/or llama3 to read a user's discord messages and respond to them accordingly
The following features have been implemented:
	1. llama3 speaks and user can chat with them
	2. llama3 remembers previous messages
	3. User can exit the chat
*/

func run(){
	//Initialize the program with the ai model
	llm, err := ollama.New(ollama.WithModel("llama3"))
	errc(err)

	//Reader for user input
	reader := bufio.NewReader(os.Stdin)

	//Opens(creates if not existing) an sql .db file for chat memory
	db, err := sql.Open("sqlite3", "chatHistory.db")
	errc(err)

	//Initializes Ai with chat history
	chatHistory := sqlite3.NewSqliteChatMessageHistory(
		sqlite3.WithSession("chat"),
		sqlite3.WithDB(db),
	)
	conversationBuffer := memory.NewConversationBuffer(memory.WithChatHistory(chatHistory))
	llmChain := chains.NewConversation(llm, conversationBuffer)
	ctx := context.Background()

	//Loops user input until "exit" is typed
	for {
		fmt.Print("\n\nAI waiting for input: ")

		prompt, _ := reader.ReadString('\n')
		prompt = strings.TrimSpace(prompt)
		if(prompt == "exit"){
			break
		}
		fmt.Println("\n")
		//Initializes the Ai's response
		_, err := chains.Run(ctx, 
			llmChain, 
			prompt,
			//Prints message in real time.
			chains.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
				//Instead of printing, this needs to go to the chat box
				fmt.Print(string(chunk))
				return nil
			}),
		)
		errc(err)
		fmt.Println("\n")
	}
}

//Error catcher(For redundancy)
func errc(err error) {
	if err != nil {
		log.Fatal(err)
	}
}