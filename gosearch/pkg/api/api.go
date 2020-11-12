package api

// HTTP REST API сервера GoSearch.
// Прикладной интерфейс разработки для веб-приложения и других клиентов.

import (
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
	"net/http/pprof"
	"time"

	"gosearch/pkg/engine"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Service - служба API.
type Service struct {
	router *mux.Router
	engine *engine.Service
}

// ErrBadRequest - неверный запрос.
var ErrBadRequest = errors.New("неверный запрос")

// Счётчики Prometheus.
var (
	searchRequestsTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "search_requests_total",
		Help: "Количество поисковых запросов.",
	})
	searchRequestsTime = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "search_request_time",
		Help:    "Время выполнения поискового запроса, мс.",
		Buckets: prometheus.LinearBuckets(10, 10, 20), // 20 корзин, начиная с 0, по 10 элементов
	})
)

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
	// метрики Prometheus
	s.router.Handle("/metrics", promhttp.Handler())
	// профилирование приложения с помощьюpprof
	s.router.HandleFunc("/debug/pprof/", pprof.Index)
	s.router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	s.router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	s.router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	s.router.HandleFunc("/debug/pprof/trace", pprof.Trace)
	s.router.Handle("/debug/pprof/allocs", pprof.Handler("allocs"))
	s.router.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	s.router.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	s.router.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	s.router.Handle("/debug/pprof/block", pprof.Handler("block"))
	// метрики Prometheus
	s.router.Handle("/metrics", promhttp.Handler())
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

	t := time.Now()

	result := s.engine.Search(mux.Vars(r)["query"])

	// увеличение счётчика поисковых запросов
	searchRequestsTotal.Inc()
	time.Sleep(time.Duration(rand.Intn(50)) * 1000 * 1000)

	dur := time.Since(t).Milliseconds()
	searchRequestsTime.Observe(float64(dur))

	json.NewEncoder(w).Encode(result)
}
