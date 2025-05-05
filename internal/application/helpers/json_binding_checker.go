package helpers

import (
	"backend/internal/application/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleBindingError(c *gin.Context, err error) bool {
	if err != nil {
		logger.Log.Errorf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid data"})
		return true
	}
	return false
}