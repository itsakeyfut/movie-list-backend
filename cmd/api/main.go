package main

import (
	"backend/config"
	"backend/domain/movie"
	"backend/handlers"
	"backend/infrastructure/database"
	"backend/infrastructure/persistence"
	"backend/routes"
	"backend/usecase"
	"log"

	"github.com/gin-contrib/cors"
)

func main() {
	if err := config.LoadEnv(); err != nil {
		log.Fatalf("Failed to load env: %v", err)
	}

	db := database.Init()

	movieRepo := persistence.NewMovieRepository(db)
	movieService := movie.NewMovieService(movieRepo)
	movieUsecase := usecase.NewMovieUsecase(movieService)
	movieHandler := handlers.NewMovieHandler(movieUsecase)

	r := routes.SetupRouter(movieHandler)

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{config.Env.FrontendUrl}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	r.Use(cors.New(corsConfig))

	if err := r.Run(":" + config.Env.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}