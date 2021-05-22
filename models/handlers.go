package models

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func tempSearch(city string, chatID int64, userName string) (tgbotapi.PhotoConfig, error) {
	res, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + os.Getenv("apiToken"))

	if res.StatusCode == http.StatusNotFound {
		msg := tgbotapi.NewMessage(chatID, fmt.Sprint("City not found"))
		bot.Send(msg)
		return tgbotapi.PhotoConfig{}, err
	}

	wthr := Weather{}
	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &wthr)
	if err != nil {
		log.Fatal(err)
	}
	err = AddCitySearch(city, userName)
	if err != nil {
		log.Fatal(err)
	}

	image := wthr.GetImage()
	uploadPhoto := tgbotapi.NewPhotoUpload(chatID, image)

	uploadPhoto.Caption = fmt.Sprint("Temp: ", int(wthr.Main.Temp-272), " degrees", "\n", "Feels like: ", int(wthr.Main.FeelsLike), "\n", "Main: ", wthr.Weather[0].Main)
	return uploadPhoto, err
}
