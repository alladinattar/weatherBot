package telegram

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" //sqlite
)

var db *sql.DB

//func (env *Env) booksIndex(w http.ResponseWriter, r *http.Request) {
/*func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.WithFields(log.Fields{
			"package":  "telegram",
			"function": "initDB",
			"error":    err,
		}).Panic("Failed to open database: ", err)
	}

}*/

func AddCitySearch(db *sql.DB, city string, userName string) error {
	statement, _ :=
		db.Prepare("INSERT INTO users (name, citySearch) VALUES (?, ?)")
	_, err := statement.Exec(userName, city)

	return err
}

func GetHistoryByName(db *sql.DB, userName string) []string {
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
