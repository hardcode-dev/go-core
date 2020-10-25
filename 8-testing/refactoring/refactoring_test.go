package refactoring

import "testing"

func Test_rublesToUSD(t *testing.T) {
	const rubles = 1000
	res, err := rublesToUSD(rubles)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%d рублей - это %d долларов\n", rubles, res)
}
