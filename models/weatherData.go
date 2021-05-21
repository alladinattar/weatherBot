package models

import (
	"errors"
	"log"
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
	var goodWthr = "./images/goodWeather.png"
	var badWthr = "./images/badWeather.png"
	season := getSeason()
	if season == "summer" {
		if w.Main.Temp > 20 {
			return goodWthr
		}
		return badWthr

	} else if season == "autumn" {
		if w.Main.Temp > 5 {
			return goodWthr
		}
		return badWthr
	} else if season == "winter" {
		if w.Main.Temp > -10 {
			return goodWthr
		}
		return badWthr

	} else if season == "spring" {
		if w.Main.Temp > 10 {
			return goodWthr
		}
		return badWthr

	} else {
		log.Println(errors.New("error when get season"))
	}
	return ""
}
