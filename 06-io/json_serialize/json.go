package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type guitar struct {
	Manufacturer           string
	Model                  string
	IsAcoustic             bool
	NumOfStrings           int
	CoutriesOfManufacturig []string
}

func main() {
	g := guitar{
		Manufacturer:           "Fender",
		Model:                  "Stratocaster",
		IsAcoustic:             false,
		NumOfStrings:           6,
		CoutriesOfManufacturig: []string{"USA", "Indonesia"},
	}

	fmt.Printf("Оригинальный объект:\n%+v\n", g)

	bytes, err := json.Marshal(g)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("JSON-представление:\n%s\n", string(bytes))

	{
		_ = 0
	}
}
