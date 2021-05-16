package main

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/tgBot/models"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	models.InitDB("./weatherData.db")

	bot, err := tgbotapi.NewBotAPI(os.Getenv("botToken"))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	log.Println("Authorized on account ", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		user := models.User{UserName: update.Message.From.UserName}
		command := update.Message.Command()
		if len(command) != 0 {
			if command == "start" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprint("Please, enter city"))
				bot.Send(msg)
				continue
			} else if command == "history" {
				var cities string
				history := user.GetHistory()
				for _, city := range history {
					cities = cities + city + "\n"
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, cities)
				bot.Send(msg)
				continue
			}
		}
		res, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + update.Message.Text + "&appid=" + os.Getenv("apiToken"))

		if res.StatusCode == http.StatusNotFound {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprint("City not found"))
			bot.Send(msg)
			continue
		}

		if err != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprint(err))
			bot.Send(msg)
			log.Fatal(err)
		}
		wthr := models.Weather{}
		body, _ := ioutil.ReadAll(res.Body)
		err = json.Unmarshal(body, &wthr)
		if err != nil {
			log.Fatal(err)
		}
		err = user.AddCitySearch(update.Message.Text)
		if err != nil {
			log.Fatal(err)
		}

		image := wthr.GetImage()
		uploadPhoto := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, image)

		uploadPhoto.Caption = fmt.Sprint(int(wthr.Main.Temp-272), " degrees")
		bot.Send(uploadPhoto)
	}
}
