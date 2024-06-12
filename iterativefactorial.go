package piscine

func IterativeFactorial(nb int) int {
	var x int = 1

	if nb > 0 && nb <= 20 {
		for i := 1; i <= nb; i++ {
			x = x * i
		}
	} else if nb == 0 {
		return 1
	} else {
		return 0
	}
	return x
}
