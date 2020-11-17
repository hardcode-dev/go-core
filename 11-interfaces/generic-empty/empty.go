package methods

type car struct {
	model string
	price int
	year  int
}

type bike struct {
	model string
	price int
	year  int
}

func pricesSum(products ...interface{}) int {
	var sum int
	for _, p := range products {
		if p, ok := p.(car); ok {
			sum += p.price
		}
		if p, ok := p.(bike); ok {
			sum += p.price
		}
	}
	return sum
}
