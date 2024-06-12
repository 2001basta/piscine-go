package piscine

func MakeRange(min, max int) []int {
	var x []int
	if min >= max {
		return []int(nil)
	} else {
		var y int = (max - min)
		x = make([]int, y)
		j := 0
		for j < y {
			for i := min; i < max; i++ {
				x[j] = i
				j++
			}
		}
	}
	return x
}
