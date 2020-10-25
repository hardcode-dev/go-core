package table

func reverse(s string) string {
	var bytes []byte
	for i := len(s) - 1; i >= 0; i-- {
		bytes = append(bytes, s[i])
	}
	return string(bytes)
}
