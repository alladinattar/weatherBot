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
	for update := range updates {
		if update.Message == nil {
			continue
		}
		command := update.Message.Command()
		switch command {
		case "start":
			msg := StartCommand(update.Message.Chat.ID)
			_, err = b.bot.Send(msg)
			if err != nil {
				log.WithFields(log.Fields{
					"package":  "models",
					"function": "initBot",
					"error":    err,
				}).Error("Cannot send start command")
			}
		case "history":
			history := HistoryCommand(update.Message.From.UserName, update.Message.Chat.ID)
			_, err = b.bot.Send(history)
			if err != nil {
				log.WithFields(log.Fields{
					"package":  "models",
					"function": "initBot",
					"error":    err,
				}).Error("Cannot send history command")
			}
			continue
		}

		city := update.Message.Text
		if update.Message.Location != nil {
			var location LocationInfo
			city = location.GetCityByCoordinates(update.Message.Location.Latitude, update.Message.Location.Longitude)
		}

		if strings.Contains(update.Message.Text, "\n") {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Invalid input")
			_, err = b.bot.Send(msg)
			if err != nil {
				log.WithFields(log.Fields{

					"package":  "models",
					"function": "initBot",
					"error":    err,
				}).Error("Invalid input")
			}
			continue
		}
		var tempapi tempApi
		caption, image := tempapi.SearchTemp(city)
		uploadPhoto := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, image)
		uploadPhoto.Caption = caption
		if err != nil {
			log.WithFields(log.Fields{
				"package":  "models",
				"function": "initBot",
				"error":    err,
			}).Error("Cannot get photo")
		}
		err = AddCitySearch(city, update.Message.Chat.UserName)
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
