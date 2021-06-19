package models

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	time2 "time"
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
	Sys struct {
		Sunrise int `json:"sunrise"`
		Sunset  int `json:"sunset"`
	} `json:"sys"`
}

func (w Weather) getImage() string {
	image := "./images/"
	var timeOfDay string
	fmt.Println(time2.Now().Unix())
	if int(time2.Now().Unix()) < w.Sys.Sunrise || int(time2.Now().Unix()) > w.Sys.Sunset {
		timeOfDay = "night"
	} else {
		timeOfDay = "day"
	}
	switch weather := w.Weather[0].Main; weather {
	case "Clouds":
		return image + timeOfDay + "_clouds.jpg"
	case "Snow":
		return image + timeOfDay + "_snow.jpg"
	case "Rain":
		return image + timeOfDay + "_rain.jpg"
	case "Clear":
		return image + timeOfDay + "_clear.jpg"
	case "Haze":
		return image + "haze.jpg"
	default:
		log.WithFields(log.Fields{
			"package":  "models",
			"function": "GetImage",
		}).Warning("Cannot get image, unknown weather.Main")
		return ""
	}
}

func (w Weather) getSeason() string {
	month := time2.Now().Month()
	if month >= 6 && month <= 8 {
		log.WithFields(log.Fields{
			"package":  "models",
			"function": "getSeason",
		}).Info("Season was received, summer")
		return "summer"
	} else if month > 8 && month <= 11 {
		log.WithFields(log.Fields{
			"package":  "models",
			"function": "getSeason",
		}).Info("Season was received, autumn")
		return "autumn"
	} else if month > 11 && month <= 2 {
		log.WithFields(log.Fields{
			"package":  "models",
			"function": "getSeason",
		}).Info("Season was received, winter")
		return "winter"
	} else {
		log.WithFields(log.Fields{
			"package":  "models",
			"function": "getSeason",
		}).Info("Season was received, spring")
		return "spring"
	}

}
