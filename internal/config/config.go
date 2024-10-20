package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                    string
	DatabaseURL             string
	WeatherThirdPartyAPIURL string
	Environment             string
}

func LoadConfig() (*Config, error) {
	// Load .env file
	if err := godotenv.Load("../../.env"); err != nil {
		return nil, err
	}

	return &Config{
		Port:                    os.Getenv("PORT"),
		DatabaseURL:             os.Getenv("DATABASE_URL"),
		WeatherThirdPartyAPIURL: os.Getenv("WEATHER_THIRDPARTY_API_URL"),
		Environment:             os.Getenv("ENVIRONMENT"),
	}, nil
}
