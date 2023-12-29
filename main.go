package main

// * Example usage of: chatgpt.ChatCompletion * //

import (
	"fmt"
	"main/chatgpt"
)

const (
	apiKey = "-- API KEY HERE --"
)

func main() {

	cc := chatgpt.NewChatCompletion(
		apiKey,
		chatgpt.ModelGPT4,
		100,
		chatgpt.HTTPRequestHandler,
		&chatgpt.JsonMarshalHandler{},
	)

	cc.AddMessage(chatgpt.ChatGPTRoleAssistant, "Olet avustaja ja vastaat nätisti")
	cc.AddMessage(chatgpt.ChatGPTRoleUser, "Missä Muumipeikko asuu?")

	response, err := cc.HandleRequest()
	if err != nil {
		fmt.Println("Could not handle request:", err)
		return
	}

	// Do something with response
	fmt.Println(response.Choices[0].Message.Content)
}
