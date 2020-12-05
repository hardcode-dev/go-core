package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	api    *API
	router *mux.Router
}

func main() {
	srv := new(server)
	srv.router = mux.NewRouter()
	srv.api = &API{router: srv.router}
	srv.api.Endpoints()
	http.ListenAndServe(":8080", srv.router)
}

type book struct {
	Name   string
	Author string
}

var books = []book{
	{
		Name:   "The Lord Of The Rings",
		Author: "J.R.R. Tolkien",
	},
}
