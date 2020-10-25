package table

import "testing"

func Test_reverse(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "Тест №1",
			s:    "string",
			want: "gnirts",
		},
		{
			name: "Тест №1",
			s:    "ABCdef",
			want: "fedCBD",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.s); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
