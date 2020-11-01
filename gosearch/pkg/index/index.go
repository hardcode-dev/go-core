package index

// Обратный индекс отсканированных документов.

import "gosearch/pkg/crawler"

// Interface определяет контракт службы индексирования документов.
type Interface interface {
	Add([]crawler.Document)
	Search(string) []int
}
