package routes

import (
	"backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func HealthRoute(rg *gin.RouterGroup) {
	rg.GET("/healthcheck", handlers.HealthCheck)
}