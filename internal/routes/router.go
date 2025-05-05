package routes

import (
	"backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(movieHandler *handlers.MovieHandler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	HealthRoute(api)
	MovieRoute(api, movieHandler)

	return r
}