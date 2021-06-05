package main

import (
	"github.com/tgBot/models"
)

func main() {
	models.InitDB("db/weatherData.db")
	models.InitBot()

}
