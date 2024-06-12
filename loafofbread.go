package piscine

func LoafOfBread(str string) string {
	count := 0
	a := ""
	b := ""
	if len(str) >= 5 {
		for i := 0; i < len(str); i++ {
			if count == 5 {
				a += " "
				count = 0
			} else if count != 5 && str[i] == ' ' {
				continue
			} else {
				count++
				a += string(str[i])
			}
		}
		for i := 0; i < len(a); i++ {
			if a[i] == ' ' && i == len(a)-1 {
				continue
			} else {
				b += string(a[i])
			}
		}
		return b + "\n"
	} else {
		return "\n"
	}
}
