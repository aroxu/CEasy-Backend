package models

//Config exports config
type Config struct {
	Port         string `json:"PORT"`
	Debug        bool   `json:"DEBUG"`
	DefaultLimit int    `json:"DefaultLimit"`
	MaxLimit     int    `json:"MaxLimit"`
	DBHost       string `json:"DBHost"`
	DBPort       string `json:"DBPort"`
	DBUser       string `json:"DBUser"`
	DBPass       string `json:"DBPass"`
	DBDatabase   string `json:"DBDatabase"`
}
