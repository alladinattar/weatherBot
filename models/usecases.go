package models

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
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

func (location LocationInfo) GetCityByCoordinates(lat float64, long float64) string {
	res, err := http.Get("http://api.positionstack.com/v1/reverse?access_key=" + config.GeoToken + "&query=" + fmt.Sprintf("%f", lat) + "," + fmt.Sprintf("%f", long))
	if err != nil {
		log.Fatal(err)
		resp, err := ioutil.ReadAll(res.Body)
		err = json.Unmarshal(resp, &location)
		if err != nil {
			log.WithFields(log.Fields{
				"package":  "models",
				"function": "getCityByCoordinates",
				"error":    err,
			}).Error("Cannot get city by coordinates")
		}
		return location.Data[0].County
	} else {
		return ""
	}
}

type tempApi struct {
	weather Weather
}

func (t tempApi) SearchTemp(city string) (string, string) {
	res, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + config.ApiToken)
	if res.StatusCode == http.StatusNotFound {
		return "City not Found", ""
	}
	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &t.weather)
	if err != nil {
		log.Fatal(err)
	}
	return "Temp: " + strconv.Itoa(int(t.weather.Main.Temp-272)) +
		" C°\n" + "Feels like: " + strconv.Itoa(int(t.weather.Main.FeelsLike-272)) +
		" C°\n" + "Main: " + t.weather.Weather[0].Main, t.weather.getImage()
}
