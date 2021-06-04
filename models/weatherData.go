package models

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

type Weather struct {
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
	} `json:"main"`
}

func (w Weather) GetImage() string {
	var clouds = "./images/clouds.jpg"
	var snow = "./images/snow.jpg"
	var rain = "./images/rain.jpg"
	var clear = "./images/clear.jpg"
	season := getSeason()
	fmt.Println(season)
	if w.Weather[0].Main == "Clouds" {
		return clouds
	} else if w.Weather[0].Main == "Snow" {
		return snow
	} else if w.Weather[0].Main == "Rain" {
		return rain
	} else if w.Weather[0].Main == "Clear" {
		return clear
	} else {
		log.WithFields(log.Fields{
			"package":  "models",
			"function": "GetImage",
		}).Warning("Cannot get image, unknown weather.Main")
		return ""
	}
}
