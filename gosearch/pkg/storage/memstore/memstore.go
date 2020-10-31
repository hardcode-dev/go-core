package memstore

import (
	"gosearch/pkg/crawler"
	"sort"
	"sync"
)

// DB - хранилище данных
type DB struct {
	mux  *sync.Mutex
	docs []crawler.Document
}

// New - конструктор.
func New() *DB {
	db := DB{
		mux: new(sync.Mutex),
	}
	return &db
}

// StoreDocs обавляет новые документы.
func (db *DB) StoreDocs(docs []crawler.Document) error {
	db.docs = append(db.docs, docs...)
	sort.Slice(db.docs, func(i, j int) bool { return db.docs[i].ID > db.docs[j].ID })
	return nil
}

// Docs возвращает документы по их номерам.
func (db *DB) Docs(ids []int) []crawler.Document {
	var result []crawler.Document
	db.mux.Lock()
	defer db.mux.Unlock()
	for _, id := range ids {
		for _, d := range db.docs {
			if d.ID == id {
				result = append(result, d)
				break
			}
		}
	}
	return result
}
