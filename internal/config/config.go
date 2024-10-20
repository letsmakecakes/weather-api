package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                    string
	DatabaseURL             string
	WeatherThirdPartyAPIURL string
	APIKey                  string
	RateLimit               string
	Burst                   string
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
		APIKey:                  os.Getenv("API_KEY"),
		RateLimit:               os.Getenv("RATE_LIMIT"),
		Burst:                   os.Getenv("BURST"),
		Environment:             os.Getenv("ENVIRONMENT"),
	}, nil
}
