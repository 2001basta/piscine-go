package piscine

func Any(f func(string) bool, a []string) bool {
	var x bool
	for _, y := range a {
		x = f(y)
		if x == true {
			return true
		} else {
			continue
		}
	}
	return false
}
