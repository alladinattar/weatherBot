package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const token = "#"
const apiKey = "#"

func main() {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Println("Authorized on account ", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		res, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + update.Message.Text + "&appid=" + apiKey)

		if res.StatusCode == http.StatusNotFound {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprint("Город не найден"))
			bot.Send(msg)
			continue
		}

		if err != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprint(err))
			bot.Send(msg)
			log.Fatal(err)
		}

		response := weather{}
		body, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(body, &response)
		log.Println(response.Main.Temp)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprint(int(response.Main.Temp-272), " градусов"))
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
