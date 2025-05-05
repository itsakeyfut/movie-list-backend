package persistence

import (
	"backend/internal/domain/movie"

	"gorm.io/gorm"
)

type MovieRepositoryImpl struct {
	DB *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *MovieRepositoryImpl {
	return &MovieRepositoryImpl{DB: db}
}

func (r *MovieRepositoryImpl) Save(movie *movie.Movie) (*movie.Movie, error) {
	if err := r.DB.Create(movie).Error; err != nil {
		return nil, err
	}
	return movie, nil
}

func (r *MovieRepositoryImpl) FindById(id int64) (*movie.Movie, error) {
	var movie movie.Movie
	if err := r.DB.First(&movie, id).Error; err != nil {
		return nil, err
	}
	return &movie, nil
}