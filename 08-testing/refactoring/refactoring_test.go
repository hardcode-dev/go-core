package refactoring

import (
	"testing"
)

func Test_rublesToUSD(t *testing.T) {
	const rubles = 1000
	res, err := rublesToUSD(rubles)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%d рублей - это %d долларов\n", rubles, res)
}

func Test_calc(t *testing.T) {
	type args struct {
		data   Rates
		rubles int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calc(tt.args.data, tt.args.rubles)
			if (err != nil) != tt.wantErr {
				t.Errorf("calc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("calc() = %v, want %v", got, tt.want)
			}
		})
	}
}
