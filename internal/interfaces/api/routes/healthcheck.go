package routes

import (
	"backend/internal/interfaces/api/handlers"

	"github.com/gin-gonic/gin"
)

func HealthRoute(rg *gin.RouterGroup) {
	rg.GET("/healthcheck", handlers.HealthCheck)
}