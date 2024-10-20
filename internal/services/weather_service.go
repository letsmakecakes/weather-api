package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"weatherapi/internal/cache"
	"weatherapi/pkg/models"
)

type WeatherService interface {
	GetWeather(city string) (*models.Weather, error)
}

type weatherService struct {
	cache      cache.WeatherCache
	apiKey     string
	weatherURL string
}

func NewWeatherService(cache cache.WeatherCache, apiKey string, weatherURL string) WeatherService {
	return &weatherService{cache, apiKey, weatherURL}
}

func (w *weatherService) GetWeather(city string) (*models.Weather, error) {
	if city == "" {
		return nil, fmt.Errorf("city is required")
	}

	// check if the weather data is in cache
	val, err := w.cache.GetCache(city)
	if err == nil && val != "" {
		var cachedWeather models.Weather
		err = json.Unmarshal([]byte(val), &cachedWeather)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
		}

		return &cachedWeather, nil
	}

	// if not in cache, fetch from 3rd party API
	weatherURL := fmt.Sprintf(w.weatherURL+"%s?unitGroup=metric&include=current&key=%s&contentType=json", city, w.apiKey)

	resp, err := http.Get(weatherURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch weather data")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read weather data")
	}

	var weather models.Weather
	err = json.Unmarshal(body, &weather)

	// Cache the weather data with a 12-hour expiration
	w.cache.SetCache(city, body, 12*time.Hour)

	return &weather, nil
}
