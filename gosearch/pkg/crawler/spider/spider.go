// Package spider реализует сканер содержимого веб-сайтов.
// Пакет позволяет получить список ссылок и заголовков страниц внутри веб-сайта по его URL.
package spider

import (
	"net/http"
	"net/url"
	"path"
	"strings"

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

// Scan осуществляет рекурсивный обход ссылок сайта, указанного в URL,
// с учётом глубины перехода по ссылкам, переданной в depth.
func (s *Service) Scan(url string, depth int) (data []crawler.Document, err error) {
	pages := make(map[string]string)

	parse(url, depth, pages)

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
func parse(link string, depth int, data map[string]string) error {
	if depth == 0 {
		return nil
	}

	response, err := http.Get(link)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	page, err := html.Parse(response.Body)
	if err != nil {
		return err
	}

	data[link] = pageTitle(page)

	if depth == 1 {
		return nil
	}

	// Парсим текущую ссылку, это будет базовым URL для ссылок, найденных на странице
	bu, err := url.Parse(link)
	if err != nil {
		return err
	}

	links := pageLinks(nil, page)
	for _, l := range links {
		u, err := url.Parse(l)
		if err != nil {
			// Ошибка парсинга URL - пропускаем ссылку и продолжаем дальше
			continue
		}
		if u.IsAbs() == true {
			// Абсолютная ссылка - оставляем как есть

		} else if strings.HasPrefix(l, "//") {
			// Абсолютная ссылка вида "//foo", добавляем схему (http/https) из базового URL
			u.Scheme = bu.Scheme

		} else if strings.HasPrefix(l, "/") {
			// Относительная ссылка вида "/foo", добаввляем схему и хост из базового URL
			u.Scheme = bu.Scheme
			u.Host = bu.Host

		} else {
			// Остальные ссылки считаем относительными от текущего пути в базовом URL: "foo", "./foo" etc
			// Добавляем схему, хост и path базового URL
			// т.е. если базовый URL http://example.com/foo/test.html, а текущая ссылка "bar.html"
			// ссылка будет превращена в http://example.com/foo/bar.html
			u.Scheme = bu.Scheme
			u.Host = bu.Host
			u.Path = path.Dir(bu.Path) + "/" + path.Clean(u.Path)
		}
		// Ссылка уже отсканирована - пропускаем
		if data[u.String()] != "" {
			continue
		}
		parse(u.String(), depth-1, data)
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
