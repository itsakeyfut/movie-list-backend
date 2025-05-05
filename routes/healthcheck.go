package routes

import (
	"backend/handlers"

	"github.com/gin-gonic/gin"
)

func HealthRoute(rg *gin.RouterGroup) {
	rg.GET("/health", handlers.HealthCheck)
}