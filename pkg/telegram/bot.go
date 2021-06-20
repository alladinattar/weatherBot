package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
)

//
type Bot struct {
	bot          *tgbotapi.BotAPI
	updateConfig tgbotapi.UpdateConfig
	env          *Env
}

func NewBot() *Bot {
	bot := Bot{}
	bot.env = NewEnv()
	var err error
	bot.bot, err = tgbotapi.NewBotAPI(bot.env.BotToken)
	if err != nil {
		log.WithFields(log.Fields{
			"package":  "telegram",
			"function": "initBot",
			"error":    err,
		}).Panic("Cannot start bot")
	}
	bot.bot.Debug = false
	log.Println("Authorized on account ", bot.bot.Self.UserName)
	bot.updateConfig = tgbotapi.NewUpdate(0)
	bot.updateConfig.Timeout = 60
	return &bot
}

func (b Bot) StartBot() {
	updates, err := b.bot.GetUpdatesChan(b.updateConfig)
	if err != nil {
		log.Fatal(err)
	}
	b.handleUpdates(updates)

}

func (b Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			b.handleCommand(update.Message)
			continue
		}
		b.handleMessage(update.Message)
	}
}
