module go-core/17-api-2/cmd/server

go 1.14

require (
	github.com/gorilla/mux v1.8.0
	go-core/17-api-2/pkg/api v0.0.0
)

replace go-core/17-api-2/pkg/api v0.0.0 => ../../pkg/api
