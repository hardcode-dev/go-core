package api

import (
	"encoding/json"
	"go-core/22-extra-movies/pkg/storage"
	"net/http"

	"github.com/gorilla/mux"
)

// API - API.
type API struct {
	db     storage.Interface
	router *mux.Router
}

// New - конструктор.
func New(db storage.Interface) *API {
	api := API{
		db:     db,
		router: mux.NewRouter(),
	}
	api.endpoints()
	return &api
}

// endpoints регистрирует обработчики API.
func (api *API) endpoints() {
	api.router.HandleFunc("/movies", api.movies).Methods(http.MethodGet, http.MethodOptions)
	api.router.HandleFunc("/movies", api.newMovie).Methods(http.MethodPost, http.MethodOptions)
	api.router.HandleFunc("/movies/{id}", api.updateMovie).Methods(http.MethodPut, http.MethodOptions)
	api.router.HandleFunc("/movies/{id}", api.deleteMovie).Methods(http.MethodDelete, http.MethodOptions)
}

func (api *API) movies(w http.ResponseWriter, r *http.Request) {
	movies, err := api.db.Movies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(movies)
}

func (api *API) newMovie(w http.ResponseWriter, r *http.Request) {

}
func (api *API) updateMovie(w http.ResponseWriter, r *http.Request) {

}

func (api *API) deleteMovie(w http.ResponseWriter, r *http.Request) {

}

// Run запускает веб-сервер.
func (api *API) Run(addr string) {
	http.ListenAndServe(addr, api.router)
}
