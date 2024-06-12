package piscine

func AlphaCount(s string) int {
	var x int = 0

	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			x += 1
		} else {
			continue
		}
	}
	return x
}
