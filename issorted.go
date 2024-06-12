package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	i := 1
	var x int

	for i < len(a) {
		x1 := f(a[0], a[1])
		x = f(a[i-1], a[i])
		i++
		if (x > 0 && x1 < 0) || (x < 0 && x1 > 0) {
			return false
		} else {
			continue
		}

	}
	return true
}
