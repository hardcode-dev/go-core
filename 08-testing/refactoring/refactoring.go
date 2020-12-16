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

	data, err := rates()
	if err != nil {
		return 0, err
	}

	rubles, err = calc(data, 1000)
	if err != nil {
		return 0, err
	}

	return rubles, nil
}

func rates() (Rates, error) {
	var data Rates
	const url = "https://www.cbr-xml-daily.ru/daily_json.js"
	resp, err := http.Get(url)
	if err != nil {
		return Rates{}, err
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return Rates{}, err
	}
	return data, nil
}

func calc(data Rates, rubles int) (int, error) {
	if data.Valute.USD.Value == 0 {
		return 0, errors.New("деление на 0")
	}
	return rubles / int(data.Valute.USD.Value), nil
}
