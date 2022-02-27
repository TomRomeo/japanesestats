package util

// Wc does a character count with a provided function
func Wc(text string, fn func(rune) bool) int {
	wc := 0
	for _, ch := range text {
		if fn(ch) {
			wc++
		}
	}
	return wc
}
