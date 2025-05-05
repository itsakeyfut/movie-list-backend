package handlers

import (
	"backend/internal/application/helpers"
	"backend/internal/application/logger"
	"backend/internal/application/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	movieUsecase *usecase.MovieUsecase
}

func NewMovieHandler(usecase *usecase.MovieUsecase) *MovieHandler {
	return &MovieHandler{movieUsecase: usecase}
}

// @Summary      Create a new movie
// @Description  Add a new movie to the database
// @Tags         movies
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        request body movieRequest true "Movie details"
// @Success      200 {object} createMovieResponse
// @Failure      400 {object} errorResponse
// @Failure      401 {object} errorResponse
// @Failure      500 {object} errorResponse
// @Router       /movies [post]
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

type movieRequest struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Image       string  `json:"image"`
	ReleaseDate string  `json:"release_date" binding:"required"`
	Rating      float64 `json:"rating"`
}

type movieResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	ReleaseDate string    `json:"release_date"`
	Rating      float64   `json:"rating"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
}

type createMovieResponse struct {
	Status string        `json:"status"`
	Data   movieResponse `json:"data"`
}