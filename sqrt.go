package piscine

func Sqrt(nb int) int {
	var x int
	var y int
	if x < 0 {
		y = 0
	} else {
		for i := 0; i <= nb; i++ {
			x = i * i
			if x == nb {
				y = i
			} else {
				continue
			}
		}
	}
	return y
}
