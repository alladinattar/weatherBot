package telegram

import (
	"database/sql"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

type Env struct {
	BotToken string `json:"botToken"`
	ApiToken string `json:"apiToken"`
	GeoToken string `json:"geoToken"`
	db       *sql.DB
}

func NewEnv() *Env {
	var env Env
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
	err = json.Unmarshal(conf, &env)
	if err != nil {
		log.WithFields(log.Fields{
			"functions": "ReadConfig",
		}).Error("Invalid json")
	}

	db, err := sql.Open("sqlite3", "db/weatherData.db")
	if err != nil {
		log.Fatal(err)
	}
	env.db = db
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, citySearch TEXT)")
	_, err = statement.Exec()
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	return &env
}
