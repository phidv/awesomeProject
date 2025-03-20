package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	router "oms/internal/api"
	"oms/internal/api/handlers"
	"oms/internal/api/middleware"
	"oms/internal/infrastructure/database"
	"oms/internal/repositories"
	"oms/internal/services"
	"oms/pkg/config"
	"oms/pkg/logger"
	"oms/pkg/validators"
)

func init() {
	config.InitApp()
}

func main() {
	logger.InitLogger(config.Global.AppName)

	db := database.PostgresqlDB(context.Background())

	server := setupServer()

	// Init Repositories, Services, Handlers
	initRepository := repositories.NewRepositories(db)
	initService := services.NewAppServices(initRepository)
	initHandler := handlers.NewAppHandlers(initService)

	router.SetupRoutes(server, initHandler)

	restPort := fmt.Sprintf(":%s", config.Global.RestPort)
	logger.Logger.Info().Msg("Starting server on port " + restPort)
	if err := server.Run(restPort); err != nil {
		logger.Logger.Fatal().Err(err).Msg("Failed to start server")
	}
}

func setupServer() *gin.Engine {
	r := gin.Default()
	validators.SetupValidator()
	r.Use(middleware.LoggingMiddleware())
	return r
}
