package models

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func StartCommand(chatID int64) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatID, fmt.Sprint("Please, enter city"))
	return msg
}

func HistoryCommand(userName string, charID int64) tgbotapi.MessageConfig {
	var cities string
	history := GetHistoryByName(userName)
	for _, city := range history {
		cities = cities + city + "\n"
	}
	msg := tgbotapi.NewMessage(charID, cities)
	return msg
}
