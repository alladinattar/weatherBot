package models

type LocationInfo struct {
	Data []struct {
		Region     string `json:"region"`
		RegionCode string `json:"region_code"`
		County     string `json:"county"`
		Locality   string `json:"locality"`
	} `json:"data"`
}

type Config struct {
	BotToken string `json:"botToken"`
	ApiToken string `json:"apiToken"`
	GeoToken string `json:"geoToken"`
}
