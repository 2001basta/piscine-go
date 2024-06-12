package piscine

func TrimAtoi(s string) int {
	var x rune
	var y []rune
	for i := 0; i < len(s); i++ {
		if (s[i] >= '0' && s[i] <= '9') || s[i] == '-' {
			x = rune(s[i])
			y = append(y, x)
		} else {
			continue
		}
	}
	var a int = 1
	som := 0
	for j := 0; j < len(y); j++ {
		if y[j] == '-' && j == 0 {
			a *= -1
			continue
		} else if y[j] == '-' && j != 0 {
			continue
		} else {
			som = som*10 + int(y[j]-'0')
		}
	}
	return som * a
}
