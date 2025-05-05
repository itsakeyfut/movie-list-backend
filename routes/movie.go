package routes

import (
	"backend/handlers"

	"github.com/gin-gonic/gin"
)

func MovieRoute(rg *gin.RouterGroup, movieHandler *handlers.MovieHandler) {
	rg.POST("/", movieHandler.CreateMovie)
}