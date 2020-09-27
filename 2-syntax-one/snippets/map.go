package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	m := make(map[int]int)
	rand.Seed(time.Now().Unix())
	for i := 0; i < 50; i++ {
		n := rand.Intn(20)
		m[n]++
	}

	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Printf("Ключ %d встречался %d раз\n", k, m[k])
	}
}
