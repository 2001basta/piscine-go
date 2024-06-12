package piscine

func BasicAtoi2(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= 48 && s[i] <= 57 {
			sum = sum*10 + int(s[i]-'0')
		} else {
			return 0
		}
	}
	return sum
}
