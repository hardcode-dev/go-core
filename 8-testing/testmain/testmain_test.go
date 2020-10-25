package testmain

import (
	"os"
	"reflect"
	"testing"
)

var db *Database

func TestMain(m *testing.M) {
	db = new(Database)
	db.connString = "postgres://user:pwd@server/database"
	db.isOK = true
	os.Exit(m.Run())
}

func TestDatabase_Products(t *testing.T) {
	want := []Product{
		{
			Name:  "Компьютер белый новый",
			Price: 20_000_00,
		},
	}
	if got := db.Products(); !reflect.DeepEqual(got, want) {
		t.Errorf("Database.Products() = %v, want %v", got, want)
	}
}
