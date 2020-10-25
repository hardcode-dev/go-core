package refactoring

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Rates - курсы валют.
type Rates struct {
	Valute struct {
		USD struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"USD"`
	}
}

func rublesToUSD(rubles int) (int, error) {
	var data Rates
	var result int
	const url = "https://www.cbr-xml-daily.ru/daily_json.js"
	resp, err := http.Get(url)
	if err != nil {
		return result, err
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return result, err
	}
	if data.Valute.USD.Value == 0 {
		return result, errors.New("деление на 0")
	}
	return rubles / int(data.Valute.USD.Value), nil
}
