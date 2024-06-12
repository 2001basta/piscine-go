package piscine

func Unmatch(a []int) int {
	var y int

	i := 0
	for i < len(a) {
		y = a[i]
		count := 0
		for w := 0; w < len(a); w++ {
			if y == a[w] {
				count++
			} else {
				continue
			}
		}

		if count%2 == 1 || count == 1 {
			return y
		}
		i++

	}
	return -1
}
