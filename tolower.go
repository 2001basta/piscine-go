package piscine

func ToLower(s string) string {
	var x string
	var y string
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			y = string(s[i] + 32)
		} else {
			y = string(s[i])
		}
		x += y
	}
	return x
}
