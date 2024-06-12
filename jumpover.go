package piscine

func JumpOver(str string) string {
	var a string
	if len(str) < 3 {
		a = ""
	} else {
		j := 0
		count := 1
		for j < len(str)-1 {
			j++
			count++
			if count == 3 {
				a += string(str[j])
				count = 0
			}

		}
	}
	return a + "\n"
}
