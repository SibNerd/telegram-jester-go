package main

import (
  "fmt"
  "github.com/Syfaro/telegram-bot-api"
  "log"
  "io/ioutil"
  )

func main() {
  // Get bot token from separate file
  bot_token, err := ioutil.ReadFile("token.txt")
    if err != nil {
          log.Fatal(err)
     }
  
  // Create bot   
	bot, err := tgbotapi.NewBotAPI(string(bot_token)) 
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
	  log.Panic(err)
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}