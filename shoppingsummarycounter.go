package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	var x string
	m := make(map[string]int)
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			x += string(str[i])
		} else {
			m[x]++
			x = ""
		}
	}
	m[x]++
	return m
}
