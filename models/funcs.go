package models

import (
	"log"
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
