package models

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	time2 "time"
)

func getSeason() string {
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

func getCityByCoordinates(lat float64, long float64) string {
	res, err := http.Get("http://api.positionstack.com/v1/reverse?access_key=" + config.GeoToken + "&query=" + fmt.Sprintf("%f", lat) + "," + fmt.Sprintf("%f", long))
	if err != nil {
		log.Fatal(err)
	}
	geoTag := LocationInfo{}
	resp, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(resp, &geoTag)
	if err != nil {
		log.WithFields(log.Fields{
			"package":  "models",
			"function": "getCityByCoordinates",
			"error":    err,
		}).Error("Cannot get city by coordinates")
	}
	return geoTag.Data[0].County
}
