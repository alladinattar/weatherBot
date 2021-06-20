package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
	"github.com/tgBot/pkg/api"
	"github.com/tgBot/pkg/db"
	"strings"
)

func (b Bot) handleCommand(message *tgbotapi.Message) {
	switch message.Command() {
	case "start":
		log.Info("Handle start command")
		msg := StartCommand(message.Chat.ID)
		_, err := b.bot.Send(msg)
		if err != nil {
			log.WithFields(log.Fields{
				"package":  "telegram",
				"function": "initBot",
				"error":    err,
			}).Error("Cannot send start command")
		}
	case "history":
		log.Info("Handle history command")
		history := HistoryCommand(b.env.db, message.From.UserName, message.Chat.ID)
		_, err := b.bot.Send(history)
		if err != nil {
			log.WithFields(log.Fields{
				"package":  "telegram",
				"function": "handleCommand",
				"error":    err,
			}).Error("Cannot send history command")
		}
	default:
		log.Info("Handle unknown command")
		_, err := b.bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Unknown command"))
		if err != nil {
			log.WithFields(log.Fields{
				"package":  "telegram",
				"function": "handleCommand",
				"error":    err,
			}).Error("Cannot send message about unknown command")
		}
	}

}
func (b Bot) handleMessage(message *tgbotapi.Message) {
	city := message.Text
	if message.Location != nil {
		log.Info("Handle location user ", message.Location.Latitude, message.Location.Latitude)
		var location api.LocationInfo
		var err error
		city, err = location.GetCityByCoordinates(b.env.GeoToken, message.Location.Latitude, message.Location.Latitude)
		if err != nil {
			log.Error("Geo API down")
			msg := tgbotapi.NewMessage(message.Chat.ID, "Geo API don't work, try later")
			_, err := b.bot.Send(msg)
			if err != nil {
				log.Error("Cannot send telegram message")
			}
			return
		}
	}
	if strings.Contains(message.Text, "\n") {
		log.Error("Handle invalid message")
		msg := tgbotapi.NewMessage(message.Chat.ID, "Invalid input")
		_, err := b.bot.Send(msg)
		if err != nil {
			log.WithFields(log.Fields{

				"package":  "telegram",
				"function": "initBot",
				"error":    err,
			}).Error("Invalid input")
		}
		return
	}
	caption, image := api.SearchTemp(b.env.ApiToken, city)
	uploadPhoto := tgbotapi.NewPhotoUpload(message.Chat.ID, image)
	uploadPhoto.Caption = caption

	err := db.AddCitySearch(b.env.db, city, message.Chat.UserName)
	if err != nil {
		log.Fatal(err)
	}
	_, err = b.bot.Send(uploadPhoto)
	if err != nil {
		log.WithFields(log.Fields{
			"package":  "telegram",
			"function": "initBot",
			"error":    err,
		}).Error("Cannot send photo")
	}

}
