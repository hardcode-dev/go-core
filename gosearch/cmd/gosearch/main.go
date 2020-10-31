package main

import (
	"log"
	"net/http"

	"gosearch/pkg/api"
	"gosearch/pkg/crawler"
	"gosearch/pkg/crawler/spider"
	"gosearch/pkg/engine"
	"gosearch/pkg/index"
	"gosearch/pkg/index/hash"
	"gosearch/pkg/storage"
	"gosearch/pkg/storage/memstore"

	"github.com/gorilla/mux"
)

// Сервер Интернет-поисковика GoSearch.
type gosearch struct {
	api     *api.Service
	engine  *engine.Service
	scanner crawler.Interface
	index   index.Interface
	storage storage.Interface

	router *mux.Router

	sites []string
	depth int
	addr  string
}

func main() {
	server := new()
	server.init()
	server.run()
}

func new() *gosearch {
	gs := gosearch{}
	gs.router = mux.NewRouter()
	gs.scanner = spider.New()
	gs.index = hash.New()
	gs.storage = memstore.New()
	gs.engine = engine.New(gs.index, gs.storage)
	gs.api = api.New(gs.router, gs.engine)
	gs.sites = []string{"https://go.dev", "https://golang.org/"}
	gs.depth = 2
	gs.addr = ":80"
	return &gs
}

func (gs *gosearch) init() {
	log.Println("Сканирование сайтов.")
	id := 0
	for _, url := range gs.sites {
		log.Println("Сайт:", url)
		data, err := gs.scanner.Scan(url, gs.depth)
		if err != nil {
			continue
		}
		for i := range data {
			data[i].ID = id
			id++
		}
		log.Println("Индексирование документов.")
		gs.index.Add(data)
		log.Println("Сохранение документов.")
		err = gs.storage.StoreDocs(data)
		if err != nil {
			log.Println("ошибка при добавлении сохранении документов в БД:", err)
			continue
		}
	}
}

func (gs *gosearch) run() {
	log.Println("Запуск http-сервера на интерфейсе:", gs.addr)
	http.ListenAndServe(gs.addr, gs.router)
}
