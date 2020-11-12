module gosearch/cmd/gosearch

go 1.14

replace (
	gosearch/pkg/api => ../../pkg/api
	gosearch/pkg/crawler => ../../pkg/crawler
	gosearch/pkg/crawler/membot => ../../pkg/crawler/membot
	gosearch/pkg/crawler/spider => ../../pkg/crawler/spider
	gosearch/pkg/engine => ../../pkg/engine
	gosearch/pkg/index => ../../pkg/index
	gosearch/pkg/index/hash => ../../pkg/index/hash
	gosearch/pkg/storage => ../../pkg/storage
	gosearch/pkg/storage/memstore => ../../pkg/storage/memstore
)

require (
	github.com/gorilla/mux v1.8.0
	google.golang.org/protobuf v1.25.0 // indirect
	gosearch/pkg/api v0.0.0-00010101000000-000000000000
	gosearch/pkg/crawler v0.0.0-00010101000000-000000000000
	gosearch/pkg/crawler/spider v0.0.0-00010101000000-000000000000
	gosearch/pkg/engine v0.0.0-00010101000000-000000000000
	gosearch/pkg/index v0.0.0-00010101000000-000000000000
	gosearch/pkg/index/hash v0.0.0-00010101000000-000000000000
	gosearch/pkg/storage v0.0.0-00010101000000-000000000000
	gosearch/pkg/storage/memstore v0.0.0-00010101000000-000000000000
)
