package models

type User struct {
	UserName string
}

func (u User) AddCitySearch(city string) error {
	statement, _ :=
		db.Prepare("INSERT INTO users (name, citySearch) VALUES (?, ?)")
	_, err := statement.Exec(u.UserName, city)
	return err
}

func (u User) GetHistory() []string {
	rows, _ :=
		db.Query("SELECT citySearch FROM users WHERE name = '" + u.UserName + "'")
	var city string
	var cities []string
	for rows.Next() {
		rows.Scan(&city)
		cities = append(cities, city)
	}
	return cities
}
