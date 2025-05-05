package routes

import (
	"backend/internal/interfaces/api/handlers"

	"github.com/gin-gonic/gin"
)

func AuthRoute(rg *gin.RouterGroup, authHandler *handlers.AuthHandler) {
	auth := rg.Group("/auth")
	auth.POST("/signup", authHandler.Signup)
	auth.POST("/login", authHandler.Login)
}
