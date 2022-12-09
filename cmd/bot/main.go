package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5732583169:AAGcjnJv15W9zlWEmA4AuJg0_SmwwC5Fhxk")
	if err != nil {
		log.Panic(err)
	}
	//bot.Debug = true

	log.Printf("Account: %v", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
		Offset:  0,
		Limit:   0,
	}
	updates := bot.GetUpdatesChan(u)
	for update := range updates {

		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
	}
}
