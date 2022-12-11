package main

import (
	"bufio"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//const token = "5732583169:AAGcjnJv15W9zlWEmA4AuJg0_SmwwC5Fhxk"

func main() {
	bot, err := tgbotapi.NewBotAPI(GetToken())
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

func GetToken() string {
	file, err := os.Open("token.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var token string
	for scanner.Scan() {
		token = scanner.Text()
	}
	
	return token
}
