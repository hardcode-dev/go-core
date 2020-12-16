package storage

import (
	"os"
)

// PgDB позволяет работать с БД PostgreSQL.
type PgDB struct {
	connString string // строка подключения к БД
}

// NewPgDB - конструктор.
func NewPgDB() *PgDB {
	db := PgDB{
		connString: os.Getenv("pg_conn_string"),
	}
	return &db
}

// Users возвращает список пользователей из БД.
func (p *PgDB) Users() []User {
	// поскольку БД нет - возвращаем так
	users := []User{
		{
			Name: "Kurt Cobain",
		},
	}
	return users
}
