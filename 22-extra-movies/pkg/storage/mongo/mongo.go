package mongo

import (
	"context"
	"go-core/22-extra-movies/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	m "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB - БД Монго.
type DB struct {
	client *m.Client
}

// New - конструктор.
func New() (*DB, error) {
	db := DB{}
	opts := options.Client().ApplyURI("mongodb://ubuntu-server.northeurope.cloudapp.azure.com:27017/")
	client, err := m.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}
	db.client = client
	return &db, nil
}

// Movies возвращает все фильмы.
func (db *DB) Movies() ([]models.Movie, error) {
	moviesCol := db.client.Database("catalog").Collection("movies")
	cur, err := moviesCol.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	var movies []models.Movie
	for cur.Next(context.Background()) {
		var item models.Movie
		err := cur.Decode(&item)
		if err != nil {
			return nil, err
		}
		movies = append(movies, item)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return movies, nil
}

// ImportMovies импортирует пару фильмов.
func (db *DB) ImportMovies() error {
	data := []interface{}{
		models.Movie{
			ID:    0,
			Title: "Legends of the Fall",
		},
		models.Movie{
			ID:    1,
			Title: "Terminator",
		},
	}
	moviesCol := db.client.Database("catalog").Collection("movies")
	err := moviesCol.Drop(context.Background())
	if err != nil {
		return err
	}
	_, err = moviesCol.InsertMany(context.Background(), data)
	return err
}
