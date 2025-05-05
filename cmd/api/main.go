package main

import (
	"backend/internal/config"
	"backend/internal/domain/movie"
	"backend/internal/handlers"
	"backend/internal/infrastructure/database"
	"backend/internal/infrastructure/persistence"
	"backend/internal/logger"
	"backend/internal/middlewares"
	"backend/internal/routes"
	"backend/internal/usecase"
	"log"

	"github.com/gin-contrib/cors"
)

func main() {
	// ENV
	if err := config.LoadEnv(); err != nil {
		log.Fatalf("Failed to load env: %v", err)
	}

	// Logger
	logger.Init()
	logger.Log.Info("Logger initialized")

	// DB
	db := database.Init()

	// DI
	movieRepo := persistence.NewMovieRepository(db)
	movieService := movie.NewMovieService(movieRepo)
	movieUsecase := usecase.NewMovieUsecase(movieService)
	movieHandler := handlers.NewMovieHandler(movieUsecase)

	// Init Gin
	r := routes.SetupRouter(movieHandler)
	r.Use(middlewares.LogRequestBody())

	// CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{config.Env.FrontendUrl}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	r.Use(cors.New(corsConfig))

	// Run Server
	if err := r.Run(":" + config.Env.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}