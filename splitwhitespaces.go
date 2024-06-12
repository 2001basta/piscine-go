package piscine

func SplitWhiteSpaces(s string) []string {
	var a string
	var x []string
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			a += string(s[i])
		} else if s[i] == ' ' && s[i-1] != ' ' && i > 0 {
			x = append(x, a)
			a = ""
		} else if s[i] == ' ' && s[i-1] == ' ' && i > 0 {
			a = ""
		}
		if s[i] != ' ' && i == len(s)-1 {
			x = append(x, a)
		}
	}
	return x
}
