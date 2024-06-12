package piscine

func DescendAppendRange(max, min int) []int {
	var x []int
	if min >= max {
		return []int{}
	} else {
		for i := max; i > min; i-- {
			x = append(x, i)
		}
	}
	return x
}
