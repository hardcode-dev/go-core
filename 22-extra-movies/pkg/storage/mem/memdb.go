package mem

import (
	"go-core/22-extra-movies/pkg/models"
)

// DB - БД в памяти.
type DB int

// Movies возвращает фильмы.
func (db *DB) Movies() ([]models.Movie, error) {
	return []models.Movie{
		{
			ID:    0,
			Title: "Legends of the Fall",
		},
		{
			ID:    1,
			Title: "Terminator",
		},
	}, nil
}
