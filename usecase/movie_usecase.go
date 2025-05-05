package usecase

import "backend/domain/movie"

type MovieUsecase struct {
	movieService *movie.MovieService
}

func NewMovieUsecase(service *movie.MovieService) *MovieUsecase {
	return &MovieUsecase{movieService: service}
}

func (uc *MovieUsecase) AddMovie(title, description, releaseDate string, rating float64) (*movie.Movie, error) {
	return uc.movieService.AddMovie(title, description, releaseDate, rating)
}