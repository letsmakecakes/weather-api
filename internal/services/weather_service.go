package services

import "weatherapi/pkg/models"

type WeatherService interface {
	GetWeather(city string) (*models.Weather, error)
}
