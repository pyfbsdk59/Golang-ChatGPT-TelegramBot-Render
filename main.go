package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	gogpt "github.com/sashabaranov/go-gpt3"
	"github.com/yanzay/tbot/v2"
)

func getChatGPTresponse(ctx context.Context, question string) string {
	c := gogpt.NewClient(os.Getenv("OPENAI_TOKEN"))

	maxtokens, err0 := strconv.Atoi(os.Getenv("OPENAI_MAXTOKENS"))

	if err0 != nil {
		fmt.Println("Error during conversion")
		return "MaxTokens Conversion Error happened!"
	}

	req := gogpt.CompletionRequest{
		Model:       "text-davinci-003",
		MaxTokens:   maxtokens,
		Prompt:      question, //it represents the text your typed in
		Temperature: 0,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return "You got an error!"
	} else {
		fmt.Println(resp.Choices[0].Text)

		return resp.Choices[0].Text // it represents the answer ChatGPT gives you
	}

}

func main() {
	ctx := context.Background()
	bot := tbot.New(os.Getenv("TELEGRAM_BOT_TOKEN"),
		tbot.WithWebhook("", ":8080"))
	c := bot.Client() //Please add your Render URL between "". 請在引號中加入你的Render網址

	bot.HandleMessage(".*human:*", func(m *tbot.Message) { //The AI must be triggered by the keyword "human:"
		c.SendMessage(m.Chat.ID, "AI:"+getChatGPTresponse(ctx, m.Text)) //m.Text represents the text you typed in 代表你打的文字
	})
	log.Fatal(bot.Start())
}
