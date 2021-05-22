package models

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3" //sqlite driver
	"log"
)

var db *sql.DB

func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Panic(err)
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
	fmt.Println("SELECT * FROM users WHERE name='" + userName + "' ORDER BY id DESC LIMIT 5")
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
