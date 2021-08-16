package main

import (
  "net/http"
  "encoding/json"
  "github.com/Syfaro/telegram-bot-api"
  "github.com/icelain/jokeapi"
  "log"
  "io/ioutil"
  )

// const joke_api_one := "https://official-joke-api.appspot.com/jokes/general/random"

func getJson(url string, target interface{}) error {
  // Get nice JSON from URL response
  var myClient = &http.Client{Timeout: 10 * time.Second}
    r, err := myClient.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()

    return json.NewDecoder(r.Body).Decode(target)
}

func main() {
  // Set parameters for jokeAPI
  joke_type := "single"
  blacklist := []string{"rasist", "sexist"}
  ctgs := []string{"Misc", "Pun", "Dark"}
  
  // Get bot token from separate file
  bot_token, err := ioutil.ReadFile("token.txt")
    if err != nil {
          log.Fatal(err)
     }
  
  // Connect to API with jokes
  joke_api, err := jokeapi.New()
  
  if err != nil {
  	panic(err)
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
    
    if update.Message.IsCommand() {
      log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
      
      msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
      switch update.Message.Command() {
        case "help": 
          msg.Text("Выберите команду /joke для получения случайной шутки.")
        case "joke":
          api.SetParams(&ctgs, &blacklist, &jt)
          response := api.Fetch()
          msg.Text(response)
        default: 
          msg.Text("Боюсь, я не знаю такой команды :с")
      }
		bot.Send(msg)
	}
}