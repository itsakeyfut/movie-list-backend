package movie

import "time"

type Movie struct {
    ID          int64   `json:"id"`
    Title       string  `json:"title"`
    Description string  `json:"description"`
    ReleaseDate string  `json:"release_date"`
    Rating      float64 `json:"rating"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func NewMovie(id int64, title, description, releaseDate string, rating float64) *Movie {
    return &Movie{
        ID:          id,
        Title:       title,
        Description: description,
        ReleaseDate: releaseDate,
        Rating:      rating,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
    }
}

func (m *Movie) UpdateRating(newRating float64) {
    m.Rating = newRating
}
