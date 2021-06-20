package main

import (
	"github.com/tgBot/pkg/telegram"
)

func main() {
	bot := telegram.NewBot()
	bot.StartBot()
}
