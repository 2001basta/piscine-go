package piscine

func NRune(s string, n int) rune {
	x := []rune(s)
	if n > 0 && n <= len(x) {
		for i := 1; i <= len(x); i++ {
			if i == n {
				return x[i-1]
			} else {
				continue
			}
		}
	}
	return 0
}
