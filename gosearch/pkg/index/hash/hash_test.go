package hash

import (
	"gosearch/pkg/crawler"
	"reflect"
	"testing"
)

func TestIndex_Add(t *testing.T) {
	ind := New()
	docs := []crawler.Document{
		{
			ID:    10,
			Title: "Два Слова",
		},
		{
			ID:    20,
			Title: "And Another",
		},
	}
	ind.Add(docs)
	got := len(ind.data)
	want := 4
	if got != want {
		t.Fatalf("получили %d, ожидалось %d", got, want)
	}
}

func TestIndex_Search(t *testing.T) {
	ind := New()
	docs := []crawler.Document{
		{
			ID:    10,
			Title: "Два Слова",
		},
		{
			ID:    20,
			Title: "And Another Three",
		},
		{
			ID:    30,
			Title: "Three Tokens More",
		},
	}
	ind.Add(docs)

	tests := []struct {
		name  string
		token string
		want  []int
	}{
		{
			name:  "Тест №1",
			token: "ДВА",
			want:  []int{10},
		},
		{
			name:  "Тест №2",
			token: "THree",
			want:  []int{20, 30},
		},
		{
			name:  "Тест №3",
			token: "NotAToken",
			want:  nil,
		},
		{
			name:  "Тест №4",
			token: "three tokens",
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ind.Search(tt.token); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("получили %v, ожидалось %v", got, tt.want)
			}
		})
	}
}
