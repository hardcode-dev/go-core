module gosearch/pkg/api

go 1.14

replace gosearch/pkg/crawler => ../crawler
replace gosearch/pkg/engine => ../engine
replace gosearch/pkg/index => ../index
replace gosearch/pkg/index/hash => ../index/hash
replace gosearch/pkg/storage => ../storage

require (
	github.com/gorilla/mux v1.8.0
	gosearch/pkg/crawler v0.0.0-00010101000000-000000000000
	gosearch/pkg/engine v0.0.0-00010101000000-000000000000
)
