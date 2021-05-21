package models

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

var bot *tgbotapi.BotAPI

func InitBot() {
	var err error
	bot, err = tgbotapi.NewBotAPI(os.Getenv("botToken"))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	log.Println("Authorized on account ", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}
	for update := range updates {
		if update.Message == nil {
			continue
		}
		command := update.Message.Command()
		if len(command) != 0 {
			if command == "start" {
				msg := StartCommand(update.Message.Chat.ID)
				_, err = bot.Send(msg)
				if err != nil {
					log.Fatal(err)
				}
				continue
			} else if command == "history" {
				history := HistoryCommand(update.Message.From.UserName, update.Message.Chat.ID)
				_, err = bot.Send(history)
				continue
			}
		}
		city := update.Message.Text
		if update.Message.Location.Latitude != 0 && update.Message.Location.Longitude != 0 {
			city = getCityByCoordinates(update.Message.Location.Latitude, update.Message.Location.Longitude)
		}
		uploadPhoto, _ := tempSearch(city, update.Message.Chat.ID, update.Message.From.UserName)
		_, err = bot.Send(uploadPhoto)
	}
}
