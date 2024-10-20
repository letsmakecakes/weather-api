package controllers

import "weatherapi/internal/services"

type WeatherController struct {
	Service services.WeatherService
}

func NewWeatherController(service services.WeatherService) *WeatherController {
	return &WeatherController{service}
}
