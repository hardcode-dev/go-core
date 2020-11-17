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

func (c car) costs() int {
	return c.price
}

func (b bike) costs() int {
	return b.price
}

type worth interface {
	costs() int
}

func pricesSum(products ...worth) int {
	var sum int
	for _, p := range products {
		sum += p.costs()
	}
	return sum
}
