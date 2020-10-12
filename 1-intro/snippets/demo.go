// Package snippets - snippets.
package snippets

func fact(i int) int {
	if i == 1 {
		return 1
	}
	return i * fact(i-1)
}
