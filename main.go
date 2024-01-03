package main

// * Example usage of: chatgpt.ChatCompletion * //

import (
	"fmt"
	"main/chatgpt"
	"os"
)

var apiKey string = os.Getenv("OPENAI_API_KEY")

func main() {

	cc := chatgpt.NewChatCompletion(
		apiKey,
		chatgpt.ModelGPT4,
		100,
		chatgpt.HTTPRequestHandler,
		&chatgpt.JsonMarshalHandler{},
	)

	cc.AddMessage(chatgpt.ChatGPTRoleAssistant, "Olet avustaja ja tehtävänäsi on etsiä lähteitä, eli kirjallisuutta, artikkeleita tai muita hyviä lähteitä, joista voit saada tietoa aiheesta. Lähteitä voi olla esimerkiksi kirjastojen tietokannoissa, Google Scholarissa, tai vaikka Google-haussa. Lähteitä voi olla myös esimerkiksi lehtiartikkeleissa, kirjoissa, tai vaikka dokumenteissa. Lähteitä voi olla myös esimerkiksi lehtiartikkeleissa, kirjoissa, tai vaikka dokumenteissa. Lähteitä voi olla myös esimerkiksi lehtiartikkeleissa, kirjoissa, tai vaikka dokumenteissa.")
	cc.AddMessage(chatgpt.ChatGPTRoleUser, "Onko asia selvä?")

	response, err := cc.HandleRequest()
	if err != nil {
		fmt.Println("Chatgpt return a error:", err)
		return
	}

	// Do something with the response
	fmt.Println(response.Choices[0].Message.Content)
}
