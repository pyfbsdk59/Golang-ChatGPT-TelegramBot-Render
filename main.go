package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	gogpt "github.com/sashabaranov/go-gpt3"
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
		Prompt:      question,
		Temperature: 0,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return "You got an error!"
	} else {
		fmt.Println(resp.Choices[0].Text)

		return resp.Choices[0].Text
	}

}

/*
func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	wh, _ := tgbotapi.NewWebhookWithCert("https://www.example.com:8443/"+bot.Token, "cert.pem")

	_, err = bot.Request(wh)
	if err != nil {
		log.Fatal(err)
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}

	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "key.pem", nil)

	for update := range updates {
		log.Printf("%+v\n", update)
	}
}
*/

func main() {
	ctx := context.Background()
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			reply_msg := tgbotapi.NewMessage(update.Message.Chat.ID, getChatGPTresponse(ctx, update.Message.Text))
			//bot.Send(msg)
			bot.Send(reply_msg)
		}
	}
}
