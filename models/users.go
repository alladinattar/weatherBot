package models

func AddCitySearch(city string, userName string) error {
	statement, _ :=
		db.Prepare("INSERT INTO users (name, citySearch) VALUES (?, ?)")
	_, err := statement.Exec(userName, city)
	return err
}

func GetHistoryByName(userName string) []string {
	rows, _ :=
		db.Query("SELECT citySearch FROM users WHERE name = '" + userName + "'")
	var city string
	var cities []string
	for rows.Next() {
		rows.Scan(&city)
		cities = append(cities, city)
	}
	return cities
}
