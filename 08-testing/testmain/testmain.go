package testmain

type Database struct {
	connString string
	isOK       bool
}

type Product struct {
	Name  string
	Price int
}

func (db *Database) Products() []Product {
	data := []Product{
		{
			Name:  "Компьютер белый новый",
			Price: 20_000_00,
		},
	}
	return data
}
