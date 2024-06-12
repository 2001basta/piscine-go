package piscine

func Max(a []int) int {
	i := 1

	for i < len(a) {
		if a[i-1] > a[i] {
			a[i] = a[i-1]
		} else {
			i++
		}
	}
	return a[len(a)-1]
}
