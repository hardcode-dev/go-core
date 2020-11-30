package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// сторонний маршрутизатор из пакета Gorilla
	mux := mux.NewRouter()
	// регистрация обработчика для URL `/` в маршрутизаторе по умолчанию
	mux.HandleFunc("/{name}", mainHandler).Methods(http.MethodGet)
	// старт HTTP-сервера на порту 8080 протоколоа TCP с маршрутизатором запросов по умолчанию
	http.ListenAndServe(":8080", mux)
}

// HTTP-обработчик
func mainHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	t := template.New("main")
	t, err := t.Parse("<html><body><h2>Hi, {{.}}</h2></body></html>")
	if err != nil {
		http.Error(w, "ошибка при обработке шаблона", http.StatusInternalServerError)
		return
	}
	t.Execute(w, vars["name"])
}
