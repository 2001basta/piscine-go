package piscine

func IsPrime(nb int) bool {
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
