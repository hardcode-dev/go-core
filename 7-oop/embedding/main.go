package main

import (
	"fmt"
)

type computer struct {
	model string
	processor
}

type processor struct {
	model string
	cores int
}

func (p *processor) cpuinfo() string {
	return fmt.Sprintf("%s, %d ядер", p.model, p.cores)
}

func (c *computer) String() string {
	return fmt.Sprintf("Компьютер \"%s\".\nПроцессор \"%s\".\n", c.model, c.cpuinfo())
}

func main() {
	c := &computer{
		model: "Компьютер игровой",
		processor: processor{
			model: "Байкал",
			cores: 8,
		},
	}
	fmt.Println(c)
}
