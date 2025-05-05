package handlers

import (
	"backend/internal/helpers"
	"backend/internal/logger"
	"backend/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	movieUsecase *usecase.MovieUsecase
}

func NewMovieHandler(usecase *usecase.MovieUsecase) *MovieHandler {
	return &MovieHandler{movieUsecase: usecase}
}

func (h *MovieHandler) CreateMovie(c *gin.Context) {
	var movieRequest struct {
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Image       string  `json:"image"`
		ReleaseDate string  `json:"release_date"`
		Rating      float64 `json:"rating"`
	}

	if err := c.ShouldBindJSON(&movieRequest); helpers.HandleBindingError(c, err) {
		return
	}

	movie, err := h.movieUsecase.AddMovie(
		movieRequest.Title,
		movieRequest.Description,
		movieRequest.Image,
		movieRequest.ReleaseDate,
		movieRequest.Rating,
	)

	if err != nil {
		logger.Log.Errorf("Failed to add movie: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	logger.Log.Infof("Creating movie: %s", movieRequest.Title)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   movie,
	})
}