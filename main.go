package main

import (
  "github.com/Syfaro/telegram-bot-api"
  "github.com/icelain/jokeapi"
  
  "log"
  "os"
  )

func main() {
  // Set parameters for jokeAPI
  jt := "single"
  blacklist := []string{"rasist", "sexist"}
  ctgs := []string{"Misc", "Pun", "Dark"}
  
  // Get bot token from environmental variables
  bot_token := os.Getenv("TOKEN")
    if bot_token == "" {
          log.Fatal("Cannot get token.")
    } 
  
  // Connect to API with jokes
  joke_api := jokeapi.New()
  joke_api.SetParams(&ctgs, &blacklist, &jt)
  
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
		} else if update.Message.IsCommand() {
      log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
      
      msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
      switch update.Message.Command() {
        case "start":
          msg.Text = "Hi! I'm a joker. If you wanna read a joke, simply type /joke"
        case "help": 
          msg.Text = "Type /joke to get random joke."
        case "joke":
          response, err := joke_api.Fetch()
          if err != nil {
            log.Panic(err)
          }
          msg.Text = response.Joke[0]
        default: 
          msg.Text = "I'm afraid I don't know that command :с"
      }
		  bot.Send(msg)
		} 
	}
}