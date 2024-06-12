package piscine

func Rot14(s string) string {
	var w string
	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			if s[i] >= 'a' && s[i] <= 'z' {
				if s[i] >= 'a' && s[i] <= 'l' {
					w += string(s[i] + 14)
				} else {
					w += string(s[i] - 12)
				}
			} else {
				if s[i] >= 'A' && s[i] <= 'L' {
					w += string(s[i] + 14)
				} else {
					w += string(s[i] - 12)
				}
			}
		} else {
			w += string(s[i])
		}
	}
	return w
}
