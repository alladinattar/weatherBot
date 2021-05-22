package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	time2 "time"
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

func getCityByCoordinates(lat float64, long float64) string {
	res, err := http.Get("http://api.positionstack.com/v1/reverse?access_key=" + os.Getenv("geoToken") + "&query=" + fmt.Sprintf("%f", lat) + "," + fmt.Sprintf("%f", long))
	if err != nil {
		log.Fatal(err)
	}
	geoTag := LocationInfo{}
	resp, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(resp, &geoTag)
	if err != nil {
		log.Fatal(err)
	}
	return geoTag.Data[0].County
}
