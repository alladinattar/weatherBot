package main

import (
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tgBot/models"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	time2 "time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func getSeason() string {
	month := time2.Now().Month()
	log.Println(int(month))
	if month >= 6 && month <= 8 {
		return "summer"
	} else if month > 8 && month <= 11 {
		return "autumn"
	} else if month > 11 && month <= 2 {
		return "winter"
	} else {
		return "spring"
	}
}

func getImageAboutWeather(temp float64) string {
	var goodWthr string = "./images/goodWeather.png"
	var badWthr string = "./images/badWeather.png"
	season := getSeason()
	if season == "summer" {
		if temp > 20 {
			return goodWthr
		}
		return badWthr

	} else if season == "autumn" {
		if temp > 5 {
			return goodWthr
		}
		return badWthr
	} else if season == "winter" {
		if temp > -10 {
			return goodWthr
		}
		return badWthr

	} else if season == "spring" {
		if temp > 5 {
			return goodWthr
		}
		return badWthr

	} else {
		log.Println(errors.New("error when get season"))
	}
	return ""
}

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
				for _, city := range user.GetHistory() {
					cities = city + "\n"
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
		wthr := weather{}
		body, _ := ioutil.ReadAll(res.Body)
		err = json.Unmarshal(body, &wthr)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(getSeason())
		image := getImageAboutWeather(wthr.Main.Temp - 272)
		uploadPhoto := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, image)

		uploadPhoto.Caption = fmt.Sprint(int(wthr.Main.Temp-272), " degrees")

		log.Println(wthr.Main.Temp)
		bot.Send(uploadPhoto)
	}
}
