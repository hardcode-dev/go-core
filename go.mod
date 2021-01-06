module go-core

go 1.15

replace (
	go-core/01-intro/crawler/pkg/spider v0.0.0 => ./01-intro/crawler/pkg/spider
	go-core/17-api-2/pkg/api v0.0.0 => ./17-api-2/pkg/api
)

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/rpc v1.2.0
	github.com/gorilla/sessions v1.2.1
	github.com/gorilla/websocket v1.4.2
	github.com/haisum/rpcexample v0.0.0-20151013205443-7d034ca95162
	github.com/jackc/pgtype v1.6.2
	github.com/jackc/pgx v3.6.2+incompatible
	github.com/jackc/pgx/v4 v4.10.0
	golang.org/x/net v0.0.0-20200822124328-c89045814202
	golang.org/x/oauth2 v0.0.0-20200902213428-5d25da1a8d43
)
