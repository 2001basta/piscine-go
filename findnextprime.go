package piscine

func FindNextPrime(nb int) int {
	var y bool
	var x int
	if nb < 0 {
		return 2
	} else {
		for j := 0; j < 1000; j++ {
			x = nb + j
			y = Prime(x)
			if y == false {
				continue
			} else {
				return x
			}
		}
	}
	return 0
}

func Prime(nb int) bool {
	if nb < 2 {
		return false
	} else {
		for i := 2; i < nb; i++ {
			if nb%i == 0 {
				return false
			} else {
				continue
			}
		}
		return true
	}
}
