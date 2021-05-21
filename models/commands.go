package models

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func StartCommand(chatID int64) tgbotapi.MessageConfig {
	location := tgbotapi.NewKeyboardButton("Send my location")
	location.RequestLocation = true
	keyboard := tgbotapi.NewReplyKeyboard([]tgbotapi.KeyboardButton{location})
	msg := tgbotapi.NewMessage(chatID, "Enter city")
	msg.ReplyMarkup = keyboard
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
