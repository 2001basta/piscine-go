package piscine

func ToUpper(s string) string {
	var x string
	var y string
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			y = string(s[i] - 32)
		} else {
			y = string(s[i])
		}
		x += y
	}
	return x
}
