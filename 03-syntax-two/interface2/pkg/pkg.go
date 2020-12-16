package pkg

import (
	"fmt"
)

// Person - человек.
type Person struct {
	Name string
	Age  int
}

// Info - информация.
func (p *Person) Info() string {
	return fmt.Sprintf("%s, %d лет.", p.Name, p.Age)
}
