package piscine

func Abort(a, b, c, d, e int) int {
	slice := []int{a, b, c, d, e}

	i := 1
	for i < len(slice) {
		x := slice[i-1]
		y := slice[i]
		if y < x {
			slice[i-1] = y
			slice[i] = x
			i = 1
		} else {
			i++
			continue
		}
	}
	return slice[2]
}
