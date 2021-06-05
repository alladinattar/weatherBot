package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" //sqlite
	log "github.com/sirupsen/logrus"
)

var db *sql.DB

func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.WithFields(log.Fields{
			"package":  "models",
			"function": "initDB",
			"error":    err,
		}).Panic("Failed to open database: ", err)
	}
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, citySearch TEXT)")
	statement.Exec()
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}

func AddCitySearch(city string, userName string) error {
	statement, _ :=
		db.Prepare("INSERT INTO users (name, citySearch) VALUES (?, ?)")
	_, err := statement.Exec(userName, city)

	return err
}

func GetHistoryByName(userName string) []string {
	rows, _ :=
		db.Query("SELECT citySearch FROM users WHERE name='" + userName + "' ORDER BY id DESC LIMIT 5")
	var city string
	var cities []string
	for rows.Next() {
		rows.Scan(&city)
		cities = append(cities, city)
	}
	return cities
}
