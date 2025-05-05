package main

import (
	"backend/internal/application/logger"
	"backend/internal/application/usecase"
	"backend/internal/config"
	"backend/internal/domain/movie"
	"backend/internal/domain/user"
	"backend/internal/infrastructure/database"
	"backend/internal/infrastructure/persistence"
	"backend/internal/interfaces/api/handlers"
	"backend/internal/interfaces/api/routes"
	"log"

	"github.com/gin-contrib/cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Movie API
// @version         1.0
// @description     API for movie management application built with DDD principles
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.example.com/support
// @contact.email  support@example.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /api

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
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

	userRepo := persistence.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	userUsecase := usecase.NewUserUsecase(userService)
	authHandler := handlers.NewAuthHandler(userUsecase)

	// Init Gin
	r := routes.SetupRouter(movieHandler, authHandler)

	// Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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