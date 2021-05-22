package main

import (
	"github.com/tgBot/models"
)

func main() {
	models.InitDB("./weatherData.db")
	models.InitBot()

}
