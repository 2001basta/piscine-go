package piscine

func IsPrintable(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= 32 && s[i] <= 255 {
			continue
		} else {
			return false
		}
	}
	return true
}
