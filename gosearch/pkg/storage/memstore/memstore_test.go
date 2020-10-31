package memstore

import (
	"gosearch/pkg/crawler"
	"testing"
)

func TestDB_StoreDocs(t *testing.T) {
	db := New()
	docs := []crawler.Document{{ID: 10}, {ID: 20}}
	db.StoreDocs(docs)
	got := len(db.docs)
	want := 2
	if got != want {
		t.Fatalf("получили %d, ожидалось %d", got, want)
	}
}

func TestDB_Docs(t *testing.T) {
	db := New()
	docs := []crawler.Document{{ID: 10}, {ID: 20}}
	db.StoreDocs(docs)
	got := len(db.Docs([]int{10, 20}))
	want := 2
	if got != want {
		t.Fatalf("получили %d, ожидалось %d", got, want)
	}
}
