package api

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tgBot/pkg/models"
	"io/ioutil"
	"net/http"
	"strconv"
)

type LocationInfo struct {
	Data []struct {
		Region     string `json:"region"`
		RegionCode string `json:"region_code"`
		County     string `json:"county"`
		Locality   string `json:"locality"`
	} `json:"data"`
}

func (location LocationInfo) GetCityByCoordinates(token string, lat float64, long float64) (string, error) {
	res, err := http.Get("http://api.positionstack.com/v1/reverse?access_key=" + token + "&query=" + fmt.Sprintf("%f", lat) + "," + fmt.Sprintf("%f", long))
	if res.StatusCode == http.StatusOK {
		resp, err := ioutil.ReadAll(res.Body)
		err = json.Unmarshal(resp, &location)
		if err != nil {
			log.WithFields(log.Fields{
				"package":  "telegram",
				"function": "getCityByCoordinates",
				"error":    err,
			}).Error("Cannot get city by coordinates")
		}
		return location.Data[0].County, err
	} else {
		err = errors.New("Geo API down, try later")
		return "", err
	}
}

func SearchTemp(token string, city string) (string, string) {
	res, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + token)
	if res.StatusCode == http.StatusNotFound {
		return "City not Found", "images/fail.jpg"
	}
	weather := models.NewWeather()
	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &weather)
	if err != nil {
		log.Fatal(err)
	}
	return "Temp: " + strconv.Itoa(int(weather.Main.Temp-272)) +
		" C°\n" + "Feels like: " + strconv.Itoa(int(weather.Main.FeelsLike-272)) +
		" C°\n" + "Main: " + weather.Weather[0].Main, weather.GetImage()
}
