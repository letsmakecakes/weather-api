package controllers

import (
	"net/http"
	"weatherapi/internal/services"
	"weatherapi/pkg/utils"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type WeatherController struct {
	Service services.WeatherService
}

func NewWeatherController(service services.WeatherService) *WeatherController {
	return &WeatherController{service}
}

func (c *WeatherController) GetWeather(ctx *gin.Context) {
	city := ctx.Param("city")

	weather, err := c.Service.GetWeather(city)
	if err != nil {
		log.Errorf("error getting blog: %v", err)
		utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to retreive weather")
	} else if weather == nil {
		utils.RespondWithError(ctx, http.StatusNotFound, "Weather not found")
	}

	utils.RespondWithJSON(ctx, http.StatusOK, weather)
}
