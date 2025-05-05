package routes

import (
	"backend/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(movieHandler *handlers.MovieHandler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	HealthRoute(api)

	movies := r.Group("/movies")
	MovieRoute(movies, movieHandler)

	return r
}