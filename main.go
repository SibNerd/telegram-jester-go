package main

import (
  "github.com/Syfaro/telegram-bot-api"
  "github.com/icelain/jokeapi"
  
  "log"
  "os"
  )

func main() {
  // Set parameters for jokeAPI
  joke_type := "single"
  blacklist := []string{"rasist", "sexist"}
  ctgs := []string{"Misc", "Pun", "Dark"}
  
  // Get bot token from environmental variables
  bot_token := os.Getenv("TOKEN")
    if bot_token == "" {
          log.Fatal("Cannot get token.")
    } 
  
  // Connect to API with jokes
  joke_api:= jokeapi.New()
  
  // if err != nil {
  	// log.Panic(err)
  // }
  
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
        case "help": 
          msg.Text = "Выберите команду /joke для получения случайной шутки."
        case "joke":
          joke_api.SetParams(&ctgs, &blacklist, &jt)
          response := joke_api.Fetch()
          msg.Text = response
        default: 
          msg.Text = "Боюсь, я не знаю такой команды :с"
      }
		  bot.Send(msg)
		} 
	}
}