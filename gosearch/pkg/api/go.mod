module gosearch/pkg/api

go 1.14

replace (
	gosearch/pkg/crawler => ../crawler
	gosearch/pkg/engine => ../engine
	gosearch/pkg/index => ../index
	gosearch/pkg/index/hash => ../index/hash
	gosearch/pkg/storage => ../storage
	gosearch/pkg/storage/memstore => ../storage/memstore
)

require (
	github.com/gorilla/mux v1.8.0
	github.com/prometheus/client_golang v1.8.0
	gosearch/pkg/engine v0.0.0-00010101000000-000000000000
	gosearch/pkg/index/hash v0.0.0-00010101000000-000000000000
	gosearch/pkg/storage/memstore v0.0.0-00010101000000-000000000000
)
