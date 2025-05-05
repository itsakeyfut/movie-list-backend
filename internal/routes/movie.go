package routes

import (
	"backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func MovieRoute(rg *gin.RouterGroup, movieHandler *handlers.MovieHandler) {
	movies := rg.Group("/movies")
	movies.POST("/", movieHandler.CreateMovie)
}