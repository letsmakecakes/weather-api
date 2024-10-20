package routes

import (
	"context"
	"weatherapi/internal/cache"
	"weatherapi/internal/controllers"
	"weatherapi/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func SetupRoutes(router *gin.Engine, rdb *redis.Client, ctx context.Context, apiKey string, weatherURL string) {
	// Initialize cache, service, and controller
	weatherCache := cache.NewCache(rdb, ctx)
	weatherService := services.NewWeatherService(weatherCache, apiKey, weatherURL)
	weatherController := controllers.NewWeatherController(weatherService)

	// Group routes under /api for versioning or future scalability
	api := router.Group("/api")
	{
		weather := api.Group("/weather")
		{
			weather.GET("/:city", weatherController.GetWeather)
		}
	}
}
