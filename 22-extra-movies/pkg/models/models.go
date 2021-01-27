package models

// Movie - фильм.
type Movie struct {
	ID     int
	Title  string
	Actors []Actor
}

// Actor - актёр.
type Actor struct {
	ID        int
	FirstName string
	LastName  string
}
