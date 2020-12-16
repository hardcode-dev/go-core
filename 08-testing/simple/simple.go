package simple

import "os"

func sum(a, b int) int {
	tmp := os.Getenv("TMP")
	_ = tmp
	return a + b
}
