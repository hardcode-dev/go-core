package storage

// Interface - хранилище данных.
type Interface interface {
	Users() []User
}

// User - пользователь.
type User struct {
	Name string
}
