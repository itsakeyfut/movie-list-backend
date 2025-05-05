package movie

type MovieService struct {
	movieRepo MovieRepository
}

func NewMovieService(repo MovieRepository) *MovieService {
	return &MovieService{movieRepo: repo}
}

func (s *MovieService) AddMovie(title, description, image, releaseDate string, rating float64) (*Movie, error) {
	movie := NewMovie(0, title, description, image, releaseDate, rating)
	return s.movieRepo.Save(movie)
}