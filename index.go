package piscine

func Index(s string, toFind string) int {
	x := []rune(s)
	y := []rune(toFind)
	b := len(toFind)
	var w int
	if b <= 0 {
		return 0
	} else {
		for i := 0; i < len(s); i++ {
			if b > 1 && toFind[0] == s[i] && toFind[1] == s[i+1] {
				w = i
				c := x[w : w+(b-1)]
				for j := 0; j < len(c); j++ {
					if c[j] != y[j] {
						return 0
					} else {
						continue
					}
				}
				return w

			} else if b == 1 && toFind[0] == s[i] {
				w = i
				return w
			}
		}
		return -1
	}
}
