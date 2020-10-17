// Package spider реализует сканер содержимого веб-сайтов.

// Пакет позволяет получить список ссылок и заголовков страниц внутри веб-сайта по его URL.

package spider

import (
	"testing"
)

func TestScanSite(t *testing.T) {
	const url = "https://habr.com"
	const depth = 2
	data, err := Scan(url, depth)
	if err != nil {
		t.Fatal(err)
	}

	for k, v := range data {
		t.Logf("%s -> %s\n", k, v)
	}
}

func Test_absoluteLink(t *testing.T) {
	type args struct {
		link    string
		baseurl string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"relative link", args{
			"/index.html",
			"http://example.com",
			},
			"http://example.com/index.html",
			false,
		},
		{"absolute link", args{
			"http://example.com/index.html",
			"http://example.com",
			},
			"http://example.com/index.html",
			false,
		},
		{"external link", args{
			"http://external.com/index.html",
			"http://example.com",
		},
			"http://external.com/index.html",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := absoluteLink(tt.args.link, tt.args.baseurl)
			if (err != nil) != tt.wantErr {
				t.Errorf("absoluteLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("absoluteLink() got = %v, want %v", got, tt.want)
			}
		})
	}
}