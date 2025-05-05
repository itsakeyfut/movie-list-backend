package handlers

import (
	"backend/usecase"
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
		Title string `json:"title"`
		Description string `json:"description"`
		ReleaseDate string `json:"release_date"`
		Rating float64 `json:"rating"`
	}

	if err := c.ShouldBindJSON(&movieRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid data"})
		return
	}

	movie, err := h.movieUsecase.AddMovie(movieRequest.Title, movieRequest.Description, movieRequest.ReleaseDate, movieRequest.Rating)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": movie,
	})
}