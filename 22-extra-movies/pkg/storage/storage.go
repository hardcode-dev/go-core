package storage

import (
	"go-core/22-extra-movies/pkg/models"
)

// Interface - контракт БД.
type Interface interface {
	Movies() ([]models.Movie, error)
}
