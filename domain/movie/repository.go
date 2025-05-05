package movie

type MovieRepository interface {
	Save(movie *Movie) (*Movie, error)
	FindById(id int64) (*Movie, error)
}
