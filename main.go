package main

import (
	"github.com/tgBot/models"
)

func main() {
	models.InitDB("db/weatherData.db")
	bot := models.NewBot()
	bot.StartBot()
}
