package tdd

import "testing"

func Test_fact(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Тест №1",
			n:    1,
			want: 1,
		},
		{
			name: "Тест №2",
			n:    5,
			want: 120,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fact(tt.n); got != tt.want {
				t.Errorf("fact() = %v, want %v", got, tt.want)
			}
		})
	}
}
