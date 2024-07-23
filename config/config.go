package config

import (
	"os"
)

type Config struct {
	AppName     string
	AppPort     string
	AppVersion  string
	ExpiresIn   string
	DatabaseUri string
	GoEnv       string
	JwtSecret   string
	RedisUri    string
}

var AppConfig *Config

func LoadConfig() {
	err := LoadEnv()
	if err != nil {
		panic(err)
	}
	AppConfig = &Config{
		AppName:     os.Getenv("APP_NAME"),
		AppPort:     os.Getenv("APP_PORT"),
		AppVersion:  os.Getenv("APP_VERSION"),
		ExpiresIn:   os.Getenv("EXPIRES_IN"),
		DatabaseUri: os.Getenv("DATABASE_URI"),
		GoEnv:       os.Getenv("GO_ENV"),
		JwtSecret:   os.Getenv("JWT_SECRET"),
		RedisUri:    os.Getenv("REDIS_URI"),
	}
}
