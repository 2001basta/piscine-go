package piscine

func CountIf(f func(string) bool, tab []string) int {
	var count int = 0
	var w bool
	for _, y := range tab {
		w = f(y)
		if w == true {
			count++
		} else {
			continue
		}

	}
	return count
}
