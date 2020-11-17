package methods

import "testing"

func Test_pricesSum(t *testing.T) {
	c := car{price: 100}
	b := bike{price: 50}
	sum := pricesSum(c, b)
	if sum != 150 {
		t.Fatalf("получили %d, ожидалось %d", sum, 150)
	}
}
