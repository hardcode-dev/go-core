module go-core.course/crawler

go 1.14

require golang.org/x/net v0.0.0-20201010224723-4f7140c49acb

replace (
go-core.course/crawler/pkg/spider v0.0.0 => ./pkg/spider
)
