package piscine

func IsLower(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			continue
		} else {
			return false
		}
	}
	return true
}
