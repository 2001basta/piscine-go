package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	num := []int{}
	ss := []rune(base)
	if nbr == 0 {
		z01.PrintRune(ss[0])
	}
	if Chekbase(base) {
		x := 1
		if nbr < 0 {
			x = -1
			z01.PrintRune('-')
		}
		for nbr != 0 {
			num = append([]int{x * (nbr % len(base))}, num...)
			nbr = nbr / len(base)
		}
		for i := 0; i < len(num); i++ {
			z01.PrintRune(ss[num[i]])
		}
	} else {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
}

// ////////////////////////////////////////////////////////
func Chekbase(s string) bool {
	if len(s) < 2 {
		return false
	} else {
		for i := 0; i < len(s); i++ {
			for j := 0; j < len(s); j++ {
				if (s[i] == s[j] && i != j) || s[j] == '-' || s[j] == '+' {
					return false
				}
			}
		}
		return true
	}
}
