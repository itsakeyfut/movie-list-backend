package routes

import (
	"backend/internal/interfaces/api/handlers"
	"backend/internal/interfaces/api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter(movieHandler *handlers.MovieHandler, authHandler *handlers.AuthHandler) *gin.Engine {
	r := gin.Default()

	// No Auth required
	api := r.Group("/api")
	HealthRoute(api)
	AuthRoute(api, authHandler)

	// Auth required
	auth := api.Group("/")
	auth.Use(middlewares.JWTAuthMiddleware())
	MovieRoute(auth, movieHandler)

	return r
}