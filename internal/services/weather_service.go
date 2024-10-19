package services

import (
	"weatherapi/internal/cache"
	"weatherapi/pkg/models"
)

type WeatherService interface {
	GetWeather(city string) (*models.Weather, error)
}

type weatherService struct {
	cache cache.WeatherCache
}
