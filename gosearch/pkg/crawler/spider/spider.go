// Package spider реализует сканер содержимого веб-сайтов.
// Пакет позволяет получить список ссылок и заголовков страниц внутри веб-сайта по его URL.
package spider

import (
	"log"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/net/html"

	"gosearch/pkg/crawler"
)

// Service - служба поискового робота.
type Service struct{}

// New - констрктор службы поискового робота.
func New() *Service {
	s := Service{}
	return &s
}

// BatchScan выполняет многопоточное сканирование. Функция возвращает канал с
// отсканированными документами, и канал ошибок.
//
// Функция реализует шаблон Workers Pool для ограничения количества одновременно
// запущенных потоков сканирования.
func (s *Service) BatchScan(urls []string, depth int, workers int) (<-chan crawler.Document, <-chan error) {
	chURLs := make(chan string)          // канал входных данных (адреса сайтов)
	chOut := make(chan crawler.Document) // канал выходных данных (документов)
	chErr := make(chan error)            // канал ошибок
	var wg sync.WaitGroup
	wg.Add(workers)

	// пул рабочих потоков
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for url := range chURLs {
				data, err := s.Scan(url, depth)
				if err != nil {
					log.Println("ошибка:", err)
					chErr <- err
					return
				}
				for _, doc := range data {
					log.Println("отсканирован документ:", doc)
					chOut <- doc
				}
			}
		}()
	}
	go func() {
		wg.Wait()
		log.Println("закрываются каналы ошибок и выходных данных")
		close(chErr)
		close(chOut)
	}()

	// задания для рабочих потоков
	go func() {
		for _, url := range urls {
			chURLs <- url
		}
		log.Println("закрывается канал ссылок")
		close(chURLs)
	}()

	return chOut, chErr
}

// Scan осуществляет рекурсивный обход ссылок сайта, указанного в URL,
// с учётом глубины перехода по ссылкам, переданной в depth.
func (s *Service) Scan(url string, depth int) (data []crawler.Document, err error) {
	pages := make(map[string]string)

	parse(url, url, depth, pages)

	for url, title := range pages {
		item := crawler.Document{
			URL:   url,
			Title: title,
		}
		data = append(data, item)
	}

	return data, nil
}

// parse рекурсивно обходит ссылки на странице, переданной в url.
// Глубина рекурсии задаётся в depth.
// Каждая найденная ссылка записывается в ассоциативный массив
// data вместе с названием страницы.
func parse(url, baseurl string, depth int, data map[string]string) error {
	if depth == 0 {
		return nil
	}

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	page, err := html.Parse(response.Body)
	if err != nil {
		return err
	}

	data[url] = pageTitle(page)

	if depth == 1 {
		return nil
	}
	links := pageLinks(nil, page)
	for _, link := range links {
		link = strings.TrimSuffix(link, "/")
		// относительная ссылка
		if strings.HasPrefix(link, "/") && len(link) > 1 {
			link = baseurl + link
		}
		// ссылка уже отсканирована
		if data[link] != "" {
			continue
		}
		// ссылка содержит базовый url полностью
		if strings.HasPrefix(link, baseurl) {
			parse(link, baseurl, depth-1, data)
		}
	}

	return nil
}

// pageTitle осуществляет рекурсивный обход HTML-страницы и возвращает значение элемента <tittle>.
func pageTitle(n *html.Node) string {
	var title string
	if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
		return n.FirstChild.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		title = pageTitle(c)
		if title != "" {
			break
		}
	}
	return title
}

// pageLinks рекурсивно сканирует узлы HTML-страницы и возвращает все найденные ссылки без дубликатов.
func pageLinks(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				if !sliceContains(links, a.Val) {
					links = append(links, a.Val)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = pageLinks(links, c)
	}
	return links
}

// sliceContains возвращает true если массив содержит переданное значение
func sliceContains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
