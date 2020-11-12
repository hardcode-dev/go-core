package errors

import (
	"reflect"
	"strings"
	"testing"
)

// strToWords рзделяет строку на слова
func strToWrds(s string) (words []string) {
	var last int
	for i := 0; i < len(s); i++ {
		if s[i] == byte(' ') {
			w := s[last:i]
			words = append(words, w)
			last = i + 1
		}
		if i == len(s)-1 {
			w := s[last : i+1]
			words = append(words, w)
			break
		}

	}
	if words == nil {
		words = []string{""}
	}
	return words
}

func strToWrds2(s string) []string {
	words := strings.Split(s, " ")
	return words
}

func Test_strToWrds_strToWrds2(t *testing.T) {
	tests := []struct {
		name      string
		s         string
		wantWords []string
	}{
		{
			name:      "Тест №1",
			s:         "OneWord",
			wantWords: []string{"OneWord"},
		},
		{
			name:      "Тест №2",
			s:         "The Show Must Go On",
			wantWords: []string{"The", "Show", "Must", "Go", "On"},
		},
		{
			name:      "Тест №3",
			s:         "",
			wantWords: []string{""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotWords := strToWrds(tt.s); !reflect.DeepEqual(gotWords, tt.wantWords) {
				t.Errorf("strToWrds() = %+v, want %+v", gotWords, tt.wantWords)
			}
			if gotWords := strToWrds2(tt.s); !reflect.DeepEqual(gotWords, tt.wantWords) {
				t.Errorf("strToWrds2() = %+v, want %+v", gotWords, tt.wantWords)
			}
		})
	}
}

func Benchmark_strToWrds(b *testing.B) {
	tests := []struct {
		name      string
		s         string
		wantWords []string
	}{
		{
			name: "Тест №1",
			s:    "OneWord",
		},
		{
			name: "Тест №2",
			s:    "The Show Must Go On",
		},
		{
			name: "Тест №3",
			s:    "",
		},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				words := strToWrds(tt.s)
				if len(words) > 2 {
					b.Fatalf("len(words) > 2: %d", len(words))
				}
			}
		})
	}
}

func Benchmark_strToWrds2(b *testing.B) {
	tests := []struct {
		name      string
		s         string
		wantWords []string
	}{
		{
			name: "Тест №1",
			s:    "OneWord",
		},
		{
			name: "Тест №2",
			s:    "The Show Must Go On",
		},
		{
			name: "Тест №3",
			s:    "",
		},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				words := strToWrds2(tt.s)
				_ = words
			}
		})
	}
}

// go test -run XXX -bench . -v -benchmem
// Benchmark_strToWrds
// Benchmark_strToWrds/Тест_№1
// Benchmark_strToWrds/Тест_№1-12                  22641594                51.9 ns/op            16 B/op          1 allocs/op
// Benchmark_strToWrds/Тест_№2
// Benchmark_strToWrds/Тест_№2-12                   3343213               359 ns/op             240 B/op          4 allocs/op
// Benchmark_strToWrds/Тест_№3
// Benchmark_strToWrds/Тест_№3-12                  37226732                31.5 ns/op            16 B/op          1 allocs/op
// Benchmark_strToWrds2
// Benchmark_strToWrds2/Тест_№1
// Benchmark_strToWrds2/Тест_№1-12                 23076922                52.1 ns/op            16 B/op          1 allocs/op
// Benchmark_strToWrds2/Тест_№2
// Benchmark_strToWrds2/Тест_№2-12                  7983336               148 ns/op              80 B/op          1 allocs/op
// Benchmark_strToWrds2/Тест_№3
// Benchmark_strToWrds2/Тест_№3-12                 27900098                43.5 ns/op            16 B/op          1 allocs/op
// PASS
// ok      go-core/9-bench-debug/table     8.085s
