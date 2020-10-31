package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"gosearch/pkg/engine"

	"github.com/gorilla/mux"
)

// Service - служба API.
type Service struct {
	router *mux.Router
	engine *engine.Service
}

// ErrBadRequest - неверный запрос.
var ErrBadRequest = errors.New("неверный запрос")

// New - конструктор службы API.
func New(router *mux.Router, engine *engine.Service) *Service {
	s := Service{
		router: router,
		engine: engine,
	}
	s.endpoints()
	return &s
}

func (s *Service) endpoints() {
	// поиск
	s.router.HandleFunc("/search/{query}", s.Search)
	// веб-приложение
	s.router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./webapp"))))
}

// Search ищет документы по запросу.
func (s *Service) Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	result := s.engine.Search(mux.Vars(r)["query"])

	json.NewEncoder(w).Encode(result)
}
