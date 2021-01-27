package mem

import (
	"go-core/22-extra-movies/pkg/models"
)

// DB - БД в памяти.
type DB int

func (db *DB) Movies() ([]models.Movie, error) {
	return []api.Movie{
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
