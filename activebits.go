package piscine

func ActiveBits(n int) int {
	x := n
	count := 0
	for i := 0; i <= n; i++ {
		d := x % 2
		x = x / 2
		if d == 1 {
			count++
		}
		if x == 0 {
			break
		}
	}
	return count
}
