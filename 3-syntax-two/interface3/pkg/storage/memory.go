package storage

// MemDB  - реализация контракта БД в памяти. Заглушка для тестов.
type MemDB struct {
	connString string // строка подключения к БД
}

// Users возвращает список пользователей из БД.
func (m *MemDB) Users() []User {
	// поскольку БД нет - возвращаем так
	users := []User{
		{
			Name: "Kurt Cobain",
		},
	}
	return users
}
