package models

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

//
type Bot struct {
	bot          *tgbotapi.BotAPI
	updateConfig tgbotapi.UpdateConfig
}

var config Config

func NewBot() *Bot {
	bot := Bot{}
	log.SetLevel(log.ErrorLevel)
	ReadConfig(&config)
	var err error
	bot.bot, err = tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		log.WithFields(log.Fields{
			"package":  "models",
			"function": "initBot",
			"error":    err,
		}).Panic("Cannot start bot")
	}
	bot.bot.Debug = true
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

func (b Bot) handleCommand(message *tgbotapi.Message) {
	switch message.Text {
	case "start":
		msg := StartCommand(message.Chat.ID)
		_, err := b.bot.Send(msg)
		if err != nil {
			log.WithFields(log.Fields{
				"package":  "models",
				"function": "initBot",
				"error":    err,
			}).Error("Cannot send start command")
		}
	case "history":
		history := HistoryCommand(message.From.UserName, message.Chat.ID)
		_, err := b.bot.Send(history)
		if err != nil {
			log.WithFields(log.Fields{
				"package":  "models",
				"function": "initBot",
				"error":    err,
			}).Error("Cannot send history command")
		}
	}
}
func (b Bot) handleMessage(message *tgbotapi.Message) {
	city := message.Text
	if message.Location != nil {
		var location LocationInfo
		city = location.GetCityByCoordinates(message.Location.Latitude, message.Location.Longitude)
	}

	if strings.Contains(message.Text, "\n") {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Invalid input")
		_, err := b.bot.Send(msg)
		if err != nil {
			log.WithFields(log.Fields{

				"package":  "models",
				"function": "initBot",
				"error":    err,
			}).Error("Invalid input")
		}
	}
	var tempapi tempApi
	caption, image := tempapi.SearchTemp(city)
	uploadPhoto := tgbotapi.NewPhotoUpload(message.Chat.ID, image)
	uploadPhoto.Caption = caption

	err := AddCitySearch(city, message.Chat.UserName)
	if err != nil {
		log.Fatal(err)
	}
	_, err = b.bot.Send(uploadPhoto)
	if err != nil {
		log.WithFields(log.Fields{
			"package":  "models",
			"function": "initBot",
			"error":    err,
		}).Error("Cannot send photo")
	}

}

type Config struct {
	BotToken string `json:"botToken"`
	ApiToken string `json:"apiToken"`
	GeoToken string `json:"geoToken"`
}

func ReadConfig(config *Config) {
	file, err := os.Open(".config.json")
	if err != nil {
		log.WithFields(log.Fields{
			"functions": "ReadConfig",
		}).Error("Cannot read config file")
	}
	conf, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(conf, &config)
	if err != nil {
		log.WithFields(log.Fields{
			"functions": "ReadConfig",
		}).Error("Invalid json")
	}
}
