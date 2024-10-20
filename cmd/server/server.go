package main

import (
	"log"
	"os"
	"strconv"
	"weatherapi/internal/config"
	"weatherapi/internal/middleware"
	"weatherapi/internal/routes"
	"weatherapi/pkg/db"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	rdb, err := db.InitDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Could not connect to the Redis database: %v", err)
	}
	defer rdb.Close()

	if cfg.Environment != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(middleware.GinLogrus(logger), gin.Recovery())

	rateLimit, err := strconv.Atoi(cfg.RateLimit)
	if err != nil {
		log.Fatalf("Could not convert rate limit into integer: %v", err)
	}

	burst, err := strconv.Atoi(cfg.Burst)
	if err != nil {
		log.Fatalf("Could not convert burst into integer: %v", err)
	}

	rateLimiter := middleware.NewRateLimiter(rateLimit, burst)
	router.Use(rateLimiter.RateLimitMiddleware())

	apiURL := cfg.WeatherThirdPartyAPIURL
	apiKey := cfg.APIKey
	ctx := db.CTX
	routes.SetupRoutes(router, rdb, ctx, apiKey, apiURL)

	log.Printf("Server running on port %s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
