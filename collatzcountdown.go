package piscine

func CollatzCountdown(start int) int {
	x := start
	if start <= 0 {
		return -1
	} else {
		count := 0
		for x != 1 {
			if x%2 == 0 {
				x = x / 2
			} else {
				x = x*3 + 1
			}
			count++
		}
		return count
	}
}
